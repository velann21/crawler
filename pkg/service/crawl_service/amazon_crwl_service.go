package crawl_service

import (
	"crypto/sha1"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	aj "github.com/velann21/crawler/pkg/async_jobs"
	"hash"
	"log"
)

type CrawlAmazon interface {
  CrawlAmazonMobile()
}

type CrawlAmazonImpl struct{
	Ci CollyInterface
}

type AmazonMobileData struct {
	MobileName string `json:"mobile_name"`
	Price string `json:"price"`

}

var item string
func (collyInterface CrawlAmazonImpl) CrawlAmazonMobile(){
	kafkaProd := aj.Kafka{
		Broker:"localhost:9093",
		Topic:"crawler-data",
	}
	collyInterface.Ci.OnHTML(".sg-col-inner", func(e *colly.HTMLElement) {
		e.ForEach(".sg-row",func(_ int, divs *colly.HTMLElement) {
			divs.ForEach(".sg-row", func(_ int, divs1 *colly.HTMLElement) {
				divs1.ForEach("a",func(_ int, divs2 *colly.HTMLElement) {
					divs2.DOM.Find("span").Each(
						func(i int, s *goquery.Selection) {
							val, booll := s.Attr("class")
							if booll == true && val == "a-size-medium a-color-base a-text-normal"{
								item = s.Text()
							}
							if booll == true && val == "a-price"{
								err := aj.ProduceMessage(kafkaProd, []byte(fmt.Sprintf("%x\n", CreateHash([]byte(s.Text())))), []byte(item+"$,$"+s.Text()))
								if err != nil{
									fmt.Println(err)
								}
							}
						})
				})
			})
		})
	})
	collyInterface.Ci.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	err := collyInterface.Ci.Visit("https://www.amazon.in/s?k=Mobiles&ref=nb_sb_noss_2")
	if err != nil{
       log.Fatal("Something wrong")
	}

}

func CreateHash(byteStr []byte) []byte {
	var hashVal hash.Hash
	hashVal = sha1.New()
	hashVal.Write(byteStr)

	var bytes []byte

	bytes = hashVal.Sum(nil)
	return bytes
}



type CollyInterface interface {
	OnHTML(querySelector string, cb colly.HTMLCallback)
	OnRequest(cb colly.RequestCallback)
	Visit(string2 string)error
}

type CollyImpl struct {
    Collector colly.Collector
}


func (collyImpl *CollyImpl) OnHTML(querySelector string, cb colly.HTMLCallback){
	collyImpl.Collector.OnHTML(querySelector, cb)
}

func (collyImpl *CollyImpl) OnRequest(cb colly.RequestCallback){
	collyImpl.Collector.OnRequest(cb)
}

func (collyImpl *CollyImpl) Visit(url string) error {
	err := collyImpl.Collector.Visit(url)
	return err
}
