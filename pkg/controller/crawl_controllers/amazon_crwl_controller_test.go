package crawl_controllers

import (
	"fmt"
	"testing"

)

type AmazonServiceMock struct {

}
func (amazonServiceMock AmazonServiceMock) CrawlAmazonMobile(){
	fmt.Println("Am mocked guy")
}

func TestAmazonCrawlerController_MobileDeviceCrawlController(t *testing.T) {
	controller := AmazonCrawlerController{Ca:AmazonServiceMock{}}
	controller.MobileDeviceCrawlController()
}
