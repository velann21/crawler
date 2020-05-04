package crawl_service

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"testing"
)

type CollyMock struct {

}

func (collyMock *CollyMock) OnHTML(string2 string, ele colly.HTMLCallback){

	fmt.Println("Am mocked colly man")

}

func (collyMock *CollyMock) OnRequest(ele colly.RequestCallback){

	fmt.Println("Am mocked colly man")

}

func (collyMock *CollyMock) Visit(string2 string)error{
	return errors.New("Sm")
}


func TestCrawlAmazonImpl_CrawlAmazonMobile(t *testing.T) {
	impl := CrawlAmazonImpl{&CollyMock{}}
	impl.CrawlAmazonMobile()
}
