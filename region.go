package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type Region struct {
	FirstDist int
	DistCount int
	Name      string
	Districts map[int]*District
}

const regions_url = "http://www.cvk.gov.ua/vnd2012/wp030pt001f01=900.html"

/** Url of local cvk.gov.ua copy for testing purpose 
 *  const region_url = "http://elections/regions.html"
 **/

func GetRegions() (r []Region, err error) {
	d, err := goquery.NewDocument(regions_url)
	if err != nil {
		return nil, err
	}

	d.Find("table.t2").Last().Find("tr").First().Siblings().Each(func(j int, rs *goquery.Selection) {
		var region Region
		ars := rs.Children().First()

		region.Name, err = ars.Find("a").Html()
		if err != nil {
			fmt.Println(err)
		}

		region.Name, err = StringConvert("windows-1251", region.Name)
		if err != nil {
			fmt.Println(err)
		}

		buf, _ := ars.Siblings().Eq(0).Html()
		region.FirstDist, _ = strconv.Atoi(buf[:strings.Index(buf, " ")])

		buf, _ = ars.Siblings().Eq(1).Html()
		region.DistCount, _ = strconv.Atoi(buf)

		region.Districts = make(map[int]*District, region.DistCount)

		for i := region.FirstDist; i < region.FirstDist+region.DistCount; i++ {
			region.Districts[i], _ = NewDistrict(i)
		}

		r = append(r, region)
	})
	return r, nil
}
