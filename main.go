package main

import (
	"github.com/gocolly/colly/v2"
	cCtrl "github.com/velann21/crawler/pkg/controller/crawl_controllers"
	service "github.com/velann21/crawler/pkg/service/crawl_service"
)

type User struct {
	Name string
	LastName string
	DOB string
}
func main() {
	controller := cCtrl.AmazonCrawlerController{service.CrawlAmazonImpl{&service.CollyImpl{*colly.NewCollector()}}}
	controller.MobileDeviceCrawlController()
	//kafka := aj.Kafka{
	//	Broker:"localhost:9093",
	//	Topic:"crawler-data",
	//}
	//
	//for i:=0 ;i<=10; i++{
	//	user := User{Name:"Velan"+strconv.Itoa(i), LastName:"Nandhakumar"+strconv.Itoa(i), DOB:""}
	//	userData, err := json.Marshal(user)
	//	if err != nil{
	//		fmt.Println("")
	//	}
	//	err = aj.ProduceMessage(kafka, []byte("mobiles"+strconv.Itoa(i)), userData)
	//	if err != nil{
	//		fmt.Println(err)
	//	}
	//}

}
