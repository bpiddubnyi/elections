/*
Copyright 2012 Borys Piddubnyi <zhu@zhu.org.ua>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Region struct {
	FirstDist int
	DistCount int
	Name      string
	Districts map[int]*District
}

const regions_url = "http://www.cvk.gov.ua/pls/vnd2012/wp030?PT001F01=900"

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

		region.Name, err = StringConvert(region.Name)
		if err != nil {
			fmt.Println(err)
		}

		buf, _ := ars.Siblings().Eq(0).Html()
		region.FirstDist, _ = strconv.Atoi(buf[:strings.Index(buf, "-")])

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
