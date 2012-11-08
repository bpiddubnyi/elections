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
	"encoding/csv"
	"fmt"
	"os"
	"path"
)

func PartyMapToCsv(partyMap *map[float64]float64, partyName string, csvPath string, precision int) {
	buf := make([][]string, 2)
	for i := range buf {
		buf[i] = make([]string, len(*partyMap))
	}

	var places int
	if precision == 1 {
		places = 0
	} else {
		places = precision / 10
	}


	fName := path.Join(csvPath, partyName+".csv")
	file, err := os.OpenFile(fName, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to open file %s for writing: %v\n", fName, err)
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Failed to close file %s: %v\n", fName, err)
			panic(err)
		}
	}()

	csvFile := csv.NewWriter(file)

	i := 0
	for percent, count := range *partyMap {
		buf[0][i] = fmt.Sprintf("%.*f", places, percent)
		buf[1][i] = fmt.Sprintf("%.0f", count)
		i++
	}

	csvFile.WriteAll(buf)
}
