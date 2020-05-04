package async_jobs

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

const (
	BOOTSTRAP_SERVER = "bootstrap.servers"
	IDEMPOTENCE = "enable.idempotence"
)
type Kafka struct {
	Broker string
	Topic string
	Producer *kafka.Producer
	Term chan bool
	Done chan bool
}

func (kf *Kafka) KafkaConfig() *Kafka{
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		BOOTSTRAP_SERVER: kf.Broker,
		IDEMPOTENCE: true})

	kf.Producer = p
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return kf
}

func (kf *Kafka) WithAsync() *Kafka {
	go func() {
		doTerm := false
		for !doTerm {
			select {
			case e := <-kf.Producer.Events():
				switch ev := e.(type) {
				case *kafka.Message:
					// Message delivery report
					m := ev
					if m.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
					} else {
						fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
							*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
					}

				case kafka.Error:
					e := ev
					if e.IsFatal() {
						fmt.Printf("FATAL ERROR: %v: terminating\n", e)
					} else {
						fmt.Printf("Error: %v\n", e)
					}

				default:
					fmt.Printf("Ignored event: %s\n", ev)
				}

			case <-kf.Term:
				doTerm = true
			}
		}

		close(kf.Done)
	}()
   return kf
}

func (kf *Kafka) Push(key []byte, value []byte) error{
	err := kf.Producer.Produce(&kafka.Message{
		Key:key,
		TopicPartition: kafka.TopicPartition{Topic: &kf.Topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		return err
	}
	kf.Producer.Flush(15 * 1000)
	return nil
}


func ProduceMessage(kafkaProducer Kafka,key []byte, value []byte) error{
	err := kafkaProducer.KafkaConfig().WithAsync().Push(key, value)
	if err != nil{
		return err
	}
	return nil
}

//func KafkaProducer() {
//	// For signalling termination from main to go-routine
//	termChan := make(chan bool, 1)
//	// For signalling that termination is done from go-routine to main
//	doneChan := make(chan bool)
//
//	// Go routine for serving the events channel for delivery reports and error events.
//
//
//
//
//
//
//
//	// signal termination to go-routine
//	termChan <- true
//	// wait for go-routine to terminate
//	<-doneChan
//
//	fatalErr := p.GetFatalError()
//
//	p.Close()
//
//	// Exit application with an error (1) if there was a fatal error.
//	if fatalErr != nil {
//		os.Exit(1)
//	} else {
//		os.Exit(0)
//	}
//}
