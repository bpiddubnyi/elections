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
	"github.com/PuerkitoBio/goquery"
	"html"
	"strconv"
	"strings"
)

type Precinct struct {
	Number      int
	VotersTotal int
	VotersVoted int
	VotedPerc   float64
	Parties     map[string]float64
}

type District struct {
	Number    int
	Precincts map[int]*Precinct
}

const dist_url = "http://www.cvk.gov.ua/vnd2012/wp336pt001f01=900pf7331=%d.html"

/** Url of local gtk.gov.ua copy for testing purpose
 * const dist_url = "http://elections/dist-%d.html"
 **/

var parties []string

func NewDistrict(num int) (dist *District, err error) {
	real_dist_url := fmt.Sprintf(dist_url, num)
	d, err := goquery.NewDocument(real_dist_url)
	if err != nil {
		/*Yeah, i'm just trying to connect again. That's lame but it fucking works */
		d, err = goquery.NewDocument(real_dist_url)
		if err != nil {
			fmt.Printf("Error: failed to get page '%s' again: %v\n", real_dist_url, err)
			panic(err)
		}
	}

	header := d.Find("table.t2").Last().Find("tr").First()

	if len(parties) == 0 {
		parties = make([]string, header.Children().Size()-3)
		header.Children().Slice(3, header.Children().Size()).Each(func(i int, s *goquery.Selection) {
			buf, _ := s.Html()
			buf, _ = StringConvert("windows-1251", buf)
			buf = strings.TrimSpace(html.UnescapeString(buf))
			parties[i] = buf
		})
	} else {
		header.Children().Slice(3, header.Children().Size()).Each(func(j int, s *goquery.Selection) {
			buf, _ := s.Html()
			buf, _ = StringConvert("windows-1251", buf)
			buf = strings.TrimSpace(html.UnescapeString(buf))
			if buf != parties[j] {
				fmt.Printf("'%s' not equals '%s'(%d)\n", buf, parties[j], j)
			}
		})
	}

	precRows := header.Siblings()

	dist = new(District)

	dist.Number = num
	dist.Precincts = make(map[int]*Precinct)

	precRows.Each(func(k int, s *goquery.Selection) {
		prec := new(Precinct)

		buf, _ := s.Children().First().Children().First().Html()
		prec.Number, err = strconv.Atoi(strings.TrimSpace(buf))
		if err != nil {
			fmt.Printf("Failed to covert '%s' to int: %v", buf, err)
			panic(err)
		}

		buf, _ = s.Children().Eq(1).Html()
		prec.VotersTotal, err = strconv.Atoi(strings.TrimSpace(buf))
		if err != nil {
			fmt.Printf("Failed to covert '%s' to int: %v", buf, err)
			panic(err)
		}

		buf, _ = s.Children().Eq(2).Html()
		prec.VotersVoted, _ = strconv.Atoi(strings.TrimSpace(buf))
		if err != nil {
			fmt.Printf("Failed to covert '%s' to int: %v", buf, err)
			panic(err)
		}

		/* Currently precinct.VotedPerc is not used in calculations, 
		 *  so even if following assuming is wrong it just doesn't metter */
		if prec.VotersTotal != 0 {
			prec.VotedPerc = float64(prec.VotersTotal) / 100.0
			prec.VotedPerc = float64(prec.VotersVoted) / prec.VotedPerc
		} else {
			prec.VotedPerc = 100.0
		}

		prec.Parties = make(map[string]float64, len(parties))
		s.Children().Slice(3, s.Children().Size()).Each(func(p int, s *goquery.Selection) {
			var buf string

			if f := s.Find("span"); f.Length() > 0 {
				buf, _ = f.Html()
			} else {
				buf, _ = s.Html()
			}

			if prec.VotersVoted != 0 {
				votes, _ := strconv.Atoi(strings.TrimSpace(buf))
				prec.Parties[parties[p]] = float64(votes) / (float64(prec.VotersVoted) / 100.0)
			} else {
				prec.Parties[parties[p]] = 0
			}
		})

		dist.Precincts[prec.Number] = prec
	})

	return
}
