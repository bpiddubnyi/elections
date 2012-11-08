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
	"os"
	"path"
	"strconv"
)

const (
	defaultPrecision = 10
	defaultPrePath   = "results"
)

func main() {
	/* Flag parsing*/
	var (
		verbose, help bool
		precision     int
		savePath      string
	)

	defaultPath := path.Join(defaultPrePath, "<precision>")

	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&help, "help", false, "print this help")
	flag.IntVar(&precision, "precision", defaultPrecision, "calculation precision")
	flag.StringVar(&savePath, "path", defaultPath, "path where to save results")

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

	if savePath == defaultPath {
		savePath = path.Join(defaultPrePath, strconv.Itoa(precision))
	}

	/* Receiving the info */
	regions, err := GetRegions()
	if err != nil {
		fmt.Errorf("Failed to get regions info: %v\n", err)
		return
	}

	resultMap := make(map[string]*map[float64]float64)
	regionResultMap := make(map[string]*map[string]*map[float64]float64)

	/* Calculations */
	for _, region := range regions {
		if regionResultMap[region.Name] == nil {
			pb := make(map[string]*map[float64]float64)
			regionResultMap[region.Name] = &pb
		}

		regionPartyMap := regionResultMap[region.Name]

		for _, dist := range region.Districts {
			for _, prec := range dist.Precincts {
				/* Omitting few precincts with no voters voted */
				if prec.VotersVoted == 0 {
					continue
				}

				for party, result := range prec.Parties {
					if (*regionPartyMap)[party] == nil {
						rb := make(map[float64]float64)
						(*regionPartyMap)[party] = &rb
					}

					if resultMap[party] == nil {
						b := make(map[float64]float64)
						resultMap[party] = &b
					}

					(*resultMap[party])[Round(result, precision)]++
					(*(*regionPartyMap)[party])[Round(result, precision)]++
				}
			}
		}
	}

	/* Generate overall plots */
	for party, partyMap := range resultMap {
		/* Print out calculated data */
		if verbose {
			fmt.Printf("%s: %v\n", party, *partyMap)
		}

		if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
			fmt.Printf("Failed to create dir (%s): %v\n", savePath, err)
			return
		}

		PartyMapToPlot(partyMap, party, savePath, "[ Україна ]", precision)
		PartyMapToCsv(partyMap, party, savePath, precision)
	}

	/* Generate region plots */
	for region, regionMap := range regionResultMap {
		/* Print out calculated data */
		if verbose {
			fmt.Printf("%s:\n", region)
		}

		for party, partyMap := range *regionMap {
			/* Print out calculated data */
			if verbose {
				fmt.Printf("%s: %v\n", party, *partyMap)
			}

			regionSavePath := path.Join(savePath, region)
			if err := os.MkdirAll(regionSavePath, os.ModePerm); err != nil {
				fmt.Printf("Failed to create dir (%s): %v\n", regionSavePath, err)
				return
			}

			PartyMapToPlot(partyMap, party, regionSavePath, region, precision)
			PartyMapToCsv(partyMap, party, regionSavePath, precision)
		}
	}
}
