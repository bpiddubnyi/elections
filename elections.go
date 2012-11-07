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
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

const (
	defaultPrecision = 10
	defaultPrePath   = "plots"
)

func main() {
	/* Flag parsing*/
	var (
		verbose, help bool
		precision     int
		plotPath      string
	)

	defaultPath := path.Join(defaultPrePath, "<precision>")

	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&help, "help", false, "print this help")
	flag.IntVar(&precision, "precision", defaultPrecision, "calculation precision")
	flag.StringVar(&plotPath, "path", defaultPath, "path where to save plots")

	flag.Parse()

	if help {
		fmt.Printf("Usage: ./elections [options]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		return
	}

	if precision < 1 {
		fmt.Printf("Precision should be greater or equal to one\n")
		return
	}

	if (precision != 1) && ((precision % 10) > 0) {
		fmt.Printf("Precision should be 1 (one) or should be divisible by 10 e.g 10, 100, 1000\n")
		return
	}

	if plotPath == defaultPath {
		plotPath = path.Join(defaultPrePath, strconv.Itoa(precision))
	}

	if err := os.MkdirAll(plotPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create path for plots (%s): %v\n", plotPath, err)
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

					(*resultMap[party])[Round(result, precision)]++
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

		h := plotter.NewHistogram(xys, 100*precision)
		p.Add(h)

		fname := path.Join(plotPath, party+".png")
		if err = p.Save(8, 8, fname); err != nil {
			fmt.Printf("Failed to save plot (%s): %v\n", fname, err)
		}
	}
}
