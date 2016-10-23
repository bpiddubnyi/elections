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

type region struct {
	FirstDist int
	DistCount int
	Name      string
	Districts map[int]*district
}

const regionsURL = "http://www.cvk.gov.ua/pls/vnd2012/wp030?PT001F01=900"

func parseRegions() (regions []region, err error) {
	d, err := goquery.NewDocument(regionsURL)
	if err != nil {
		return nil, err
	}

	d.Find("table.t2").Last().Find("tr").First().Siblings().Each(func(j int, rs *goquery.Selection) {
		var reg region
		ars := rs.Children().First()

		reg.Name, err = ars.Find("a").Html()
		if err != nil {
			fmt.Println(err)
		}

		reg.Name, err = winCharsetDecoder.String(reg.Name)
		if err != nil {
			fmt.Println(err)
		}

		buf, _ := ars.Siblings().Eq(0).Html()
		reg.FirstDist, _ = strconv.Atoi(buf[:strings.Index(buf, "-")])

		buf, _ = ars.Siblings().Eq(1).Html()
		reg.DistCount, _ = strconv.Atoi(buf)

		reg.Districts = make(map[int]*district, reg.DistCount)

		for i := reg.FirstDist; i < reg.FirstDist+reg.DistCount; i++ {
			reg.Districts[i], _ = newDistrict(i)
		}

		regions = append(regions, reg)
	})
	return regions, nil
}
