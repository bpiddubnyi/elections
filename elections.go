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
	"flag"
	"fmt"
	"path"
	"strconv"
)

type config struct {
	verbose   bool
	help      bool
	precision int
	path      string
}

var cfg config

const (
	defaultPrecision = 1
	defaultPrePath   = "results"
)

func main() {
	// Flag parsing
	defaultPath := path.Join(defaultPrePath, "<precision>")

	flag.BoolVar(&cfg.verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&cfg.help, "help", false, "print this help")
	flag.IntVar(&cfg.precision, "precision", defaultPrecision, "calculation precision (decimal places)")
	flag.StringVar(&cfg.path, "path", defaultPath, "path where to save results")

	flag.Parse()

	if cfg.help {
		fmt.Printf("Usage: ./elections [options]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		return
	}

	if cfg.precision < 0 {
		fmt.Printf("Precision should be greater or equal zero\n")
		return
	}

	if cfg.path == defaultPath {
		cfg.path = path.Join(defaultPrePath, strconv.Itoa(cfg.precision))
	}

	// Receiving the info
	regions, err := parseRegions()
	if err != nil {
		fmt.Printf("Failed to get regions info: %v\n", err)
		return
	}

	resultMap := make(map[string]map[string]map[float64]float64)

	// Calculations
	countryPartyMap := make(map[string]map[float64]float64)
	resultMap["Україна"] = countryPartyMap

	for _, region := range regions {
		regionPartyMap := make(map[string]map[float64]float64)
		resultMap[region.Name] = regionPartyMap

		for _, dist := range region.Districts {
			for _, prec := range dist.Precincts {
				// Omitting few precincts with no voters voted
				if prec.VotersVoted == 0 {
					continue
				}

				for party, result := range prec.Parties {
					rpM := regionPartyMap[party]
					cpM := countryPartyMap[party]

					if rpM == nil {
						b := make(map[float64]float64)
						rpM = b
						regionPartyMap[party] = rpM
					}

					if cpM == nil {
						b := make(map[float64]float64)
						cpM = b
						countryPartyMap[party] = cpM
					}

					resultRound := round(result, cfg.precision)

					rpM[resultRound]++
					cpM[resultRound]++
				}
			}
		}
	}

	// Save results
	for region, regionMap := range resultMap {
		for party, partyMap := range regionMap {
			partyMapToPlot(partyMap, party, region, &cfg)
			partyMapToCsv(partyMap, party, region, &cfg)
		}
	}
}
