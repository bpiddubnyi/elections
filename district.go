package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type District struct {
	Number    uint
	Name      string
	Precincts []Precinct
}

const dist_url = "http://www.cvk.gov.ua/vnd2012/wp336pt001f01=900pf7331=%d.html"

func NewDistrict(num int) (dist *District, err error) {
	real_dist_url := fmt.Sprintf(dist_url, num)
	d, err:= goquery.NewDocument(real_dist_url)

	d.Find("table.t2").Last().Find("tr").First().Siblings().Each(func(j int, rs *goquery.Selection) {
		fmt.Println(rs.Html())
	})
	return nil, err
}
