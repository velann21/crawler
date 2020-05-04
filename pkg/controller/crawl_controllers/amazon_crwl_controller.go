package crawl_controllers

import service "github.com/velann21/crawler/pkg/service/crawl_service"

type AmazonCrawlerController struct {
	Ca service.CrawlAmazon
}
func (amazonCrawlerController AmazonCrawlerController) MobileDeviceCrawlController() {
	amazonCrawlerController.Ca.CrawlAmazonMobile()
}
