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

type Config struct {
	verbose   bool
	help      bool
	precision int
	path      string
}

const (
	defaultPrecision = 1
	defaultPrePath   = "results"
)

func main() {
	/* Flag parsing*/
	var config Config

	defaultPath := path.Join(defaultPrePath, "<precision>")

	flag.BoolVar(&config.verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&config.help, "help", false, "print this help")
	flag.IntVar(&config.precision, "precision", defaultPrecision, "calculation precision (decimal places)")
	flag.StringVar(&config.path, "path", defaultPath, "path where to save results")

	flag.Parse()

	if config.help {
		fmt.Printf("Usage: ./elections [options]\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		return
	}

	if config.precision < 0 {
		fmt.Printf("Precision should be greater or equal zero\n")
		return
	}

	if config.path == defaultPath {
		config.path = path.Join(defaultPrePath, strconv.Itoa(config.precision))
	}

	/* Receiving the info */
	regions, err := GetRegions()
	if err != nil {
		fmt.Errorf("Failed to get regions info: %v\n", err)
		return
	}

	resultMap := make(map[string]*map[string]*map[float64]float64)

	/* Calculations */
	countryPartyMap := make(map[string]*map[float64]float64)
	resultMap["Україна"] = &countryPartyMap

	for _, region := range regions {
		regionPartyMap := make(map[string]*map[float64]float64)
		resultMap[region.Name] = &regionPartyMap

		for _, dist := range region.Districts {
			for _, prec := range dist.Precincts {
				/* Omitting few precincts with no voters voted */
				if prec.VotersVoted == 0 {
					continue
				}

				for party, result := range prec.Parties {
					rpM := regionPartyMap[party]
					cpM := countryPartyMap[party]

					if rpM == nil {
						b := make(map[float64]float64)
						rpM = &b
						regionPartyMap[party] = rpM
					}

					if cpM == nil {
						b := make(map[float64]float64)
						cpM = &b
						countryPartyMap[party] = cpM
					}

					resultRound := Round(result, config.precision)

					(*rpM)[resultRound]++
					(*cpM)[resultRound]++
				}
			}
		}
	}

	/* Save results */
	for region, regionMap := range resultMap {
		for party, partyMap := range *regionMap {
			PartyMapToPlot(partyMap, party, region, &config)
			PartyMapToCsv(partyMap, party, region, &config)
		}
	}
}
