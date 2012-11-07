package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"flag"
	"fmt"
)

func main() {
	/* Flag parsing*/
	var verbose, help bool

	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&help, "help", false, "print this help")

	flag.Parse()

	if help {
		fmt.Printf("Usage: ./elections [options]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		return
	}

	/* Receiving the info */
	regions, err := GetRegions()
	if err != nil {
		fmt.Errorf("Failed to get regions info: %v\n", err)
		return
	}

	resultMap := make(map[string]*map[float64]float64)

	/* Calculations */
	for _, region := range regions {
		for _, dist := range region.Districts {
			for _, prec := range dist.Precincts {
				/* Omitting few precincts with no voters voted */
				if prec.VotersVoted == 0 {
					continue
				}

				for party, result := range prec.Parties {
					if partyMap := resultMap[party]; partyMap == nil {
						b := make(map[float64]float64)
						resultMap[party] = &b
					}

					(*resultMap[party])[Round(result)]++
				}
			}
		}
	}

	/* Print out calculated data */
	if verbose {
		for party, partyMapPtr := range resultMap {
			fmt.Printf("%s: %v\n", party, *partyMapPtr)
		}
	}

	/* Generate plots */
	for party, partyMap := range resultMap {
		/* Convert map to XYs */
		xys := make(plotter.XYs, len(*partyMap))
		i := 0
		for x, y := range *partyMap {
			xys[i].X = x
			xys[i].Y = y
			i++
		}

		/* Create plot */
		p, err := plot.New()
		if err != nil {
			fmt.Printf("Failed to create plot: %v\n", err)
		}

		p.Title.Text = party
		p.X.Label.Text = "Голосів за партію на дільниці(%)"
		p.Y.Label.Text = "Кількість дільниць"

		h := plotter.NewHistogram(xys, 1000)
		p.Add(h)

		p.Save(8, 8, party+".png")
	}
}
