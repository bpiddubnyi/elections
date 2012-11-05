package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"html"
	"strconv"
	"strings"
)

type Party struct {
	Name  string
	Votes uint
}

type Precinct struct {
	Number      int
	VotersTotal int
	VotersVoted int
	Parties     map[string]*Party
}

type District struct {
	Number    int
	Precincts map[int]*Precinct
}

//const dist_url = "http://www.cvk.gov.ua/vnd2012/wp336pt001f01=900pf7331=%d.html"
const dist_url = "http://elections/dist-%d.html"

var parties []string

func NewDistrict(num int) (dist *District, err error) {
	real_dist_url := fmt.Sprintf(dist_url, num)
	d, err := goquery.NewDocument(real_dist_url)
	if err != nil {
		d, err = goquery.NewDocument(real_dist_url)
		if err != nil {
			fmt.Printf("Error: failed to get page '%s' again: %v\n", real_dist_url, err)
			return nil, err
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
	fmt.Printf("  ОВО №%d, кількість дільниць: %d\n", num, precRows.Size())

	dist = new(District)

	dist.Number = num
	dist.Precincts = make(map[int]*Precinct)

	precRows.Each(func(k int, s *goquery.Selection) {
		prec := new(Precinct)

		buf, _ := s.Children().First().Children().First().Html()
		prec.Number, err = strconv.Atoi(strings.TrimSpace(buf))
		if err != nil {
			fmt.Printf("Failed to covert '%s' to int: %v", buf, err)
			return
		}


		buf, _ = s.Children().Eq(1).Html()
		prec.VotersTotal, err = strconv.Atoi(strings.TrimSpace(buf))
		if err != nil {
			fmt.Printf("Failed to covert '%s' to int: %v", buf, err)
			return
		}

		buf, _ = s.Children().Eq(2).Html()
		prec.VotersVoted, _ = strconv.Atoi(strings.TrimSpace(buf))
		if err != nil {
			fmt.Printf("Failed to covert '%s' to int: %v", buf, err)
			return
		}

		var perc float32

		if prec.VotersTotal != 0 {
			perc = float32(prec.VotersTotal/100)
			perc = float32(prec.VotersVoted)/perc
		} else {
			perc = 100.0
		}

		fmt.Printf("    Дільниця №%d: Виборців: %d, Явка: %d, Явка(%%): %f%%\n", prec.Number, prec.VotersTotal, prec.VotersVoted, perc)

		dist.Precincts[prec.Number] = prec
	})

	return
}
