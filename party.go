package main

type Party struct {
	Name  string
	Votes uint
}

func (p *Party) Init(url string) {}
