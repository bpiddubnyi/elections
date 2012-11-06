package main

import (
	"code.google.com/p/plotinum/plot"
//	"code.google.com/p/plotinum/plotter"
//	"code.google.com/p/plotinum/vg"
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
		flag.PrintDefaults()
		return
	}

	/* Receiving the info */
	regions, err := GetRegions()
	if err != nil {
		fmt.Errorf("Failed to get regions info: %v\n", err)
		return
	}

	/* Printing out the results of parsing */
	if verbose {
		for _, region := range regions {
			fmt.Printf("%s [%d-%d]:\n", region.Name, region.FirstDist, region.FirstDist+region.DistCount-1)
			for d, dist := range region.Districts {
				fmt.Printf("  ОВК-%d:\n", d)
				for _, prec := range dist.Precincts {
					fmt.Printf("    ВД-%d [%d/%d/%f%%]\n", prec.Number, prec.VotersTotal, prec.VotersVoted, prec.VotedPerc)
					for p, party := range prec.Parties {
						fmt.Printf("      %s: %d\n", p, party)
					}
				}
			}
		}
	}

    p, err := plot.New()
    if err != nil {
        fmt.Printf("Failed to create plot: %v\n", err)
    }
    p.Title.Text = "Вибори до Верховної Ради України 2012"
    p.X.Label.Text = "Голосів за партію на дільниці(%)"
    p.Y.Label.Text = "Кількість дільниць"
    p.Save(8, 8, "elections.png")
}
