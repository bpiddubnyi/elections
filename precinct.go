package main

type Precinct struct {
	Number      uint
	VotersTotal uint
	VotersVoted uint
	Parties     []Party
}

func (p *Precinct) Init(url string) {}
