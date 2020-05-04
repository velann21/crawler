package objects_manager

import "github.com/gocolly/colly/v2"


type CollyObject struct {
     Collector *colly.Collector
}


func (cob CollyObject) CreateCollyObject() *colly.Collector{
	coll := colly.NewCollector()
	cob.Collector = coll
	return coll
}

func (cob CollyObject) GetCollyObject()  *colly.Collector {
	return  cob.Collector
}
