package main

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"io/ioutil"
	"strings"
)

func StringConvert(cp string, s string) (res string, err error) {
	r, err := charset.NewReader(cp, strings.NewReader(s))
	if err != nil {
		return
	}

	rb, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	return string(rb), nil
}
