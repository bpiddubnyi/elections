package main

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"code.google.com/p/go-html-transform/h5"
	"fmt"
	"net/http"
)

type Region struct {
	Name      string
	Districts []District
}

const url string = "http://www.cvk.gov.ua/vnd2012/wp030pt001f01=900.html"

func GetRegions() (r []Region, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	tr, err := charset.NewReader("windows-1251", resp.Body)
	if err != nil {
		return nil, err
	}

	p := h5.NewParser(tr)
	err = p.Parse()
	if err != nil {
		return nil, err
	}

	t := p.Tree()
	t.Walk(func(n *h5.Node) {
		fmt.Println(n.Data())
	})

	return r, nil
}
