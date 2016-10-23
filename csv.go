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

func partyMapToCsv(m *map[float64]float64, n string, r string, c *config) {
	buf := make([][]string, 2)
	for i := range buf {
		buf[i] = make([]string, len(*m))
	}

	regionDir := path.Join(c.path, r)
	if err := os.MkdirAll(regionDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create dir (%s): %v\n", regionDir, err)
		panic(err)
	}

	fileName := path.Join(regionDir, n+".csv")
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Failed to open file %s for writing: %v\n", fileName, err)
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Failed to close file %s: %v\n", fileName, err)
			panic(err)
		}
	}()

	csvFile := csv.NewWriter(file)

	i := 0
	for percent, count := range *m {
		buf[0][i] = fmt.Sprintf("%.*f", c.precision, percent)
		buf[1][i] = fmt.Sprintf("%.0f", count)
		i++
	}

	csvFile.WriteAll(buf)

	if c.verbose {
		fmt.Printf("'%s':'%s'\n", r, n)
		csvOut := csv.NewWriter(os.Stdout)
		csvOut.WriteAll(buf)
		fmt.Printf("\n")
	}
}
