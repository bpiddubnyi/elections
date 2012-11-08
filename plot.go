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
	"fmt"
	"path"
)

func PartyMapToPlot(partyMap *map[float64]float64, partyName string, plotPath string, title string, precision int) {
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

	p.Title.Text = title + " " + partyName
	p.X.Label.Text = "Голосів за партію на дільниці(%)"
	p.Y.Label.Text = "Кількість дільниць"

	h := plotter.NewHistogram(xys, 100*precision)
	p.Add(h)

	fname := path.Join(plotPath, partyName+".png")
	if err = p.Save(8, 8, fname); err != nil {
		fmt.Printf("Failed to save plot (%s): %v\n", fname, err)
	}
}
