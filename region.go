package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type Region struct {
	Id        int
	Name      string
	Districts []District
}

const url = "http://www.cvk.gov.ua/vnd2012/wp030pt001f01=900.html"

func GetRegions() (r []Region, err error) {
	d, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	d.Find("table.t2").Last().Find("tr").First().Siblings().Each(func(j int, rs *goquery.Selection) {
		var region Region
		ars := rs.Children().First().Find("a")

		href, ex := ars.Attr("href")
		if !ex {
			fmt.Println("No href attr")
		}

		eqi := strings.LastIndex(href, "=")
		doti := strings.LastIndex(href, ".")

		region.Id, err = strconv.Atoi(href[eqi+1 : doti])
		if err != nil {
			fmt.Println(err)
		}

		region.Name, err = ars.Html()
		if err != nil {
			fmt.Println(err)
		}

		region.Name, err = StringConvert("windows-1251", region.Name)
		if err != nil {
			fmt.Println(err)
		}

		r = append(r, region)
	})
	return r, nil
}
