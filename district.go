package main

type District struct {
	Number    uint
	Name      string
	Precincts []Precinct
}

func (d *District) Init(url string) {}
