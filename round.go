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

import "math"

// Seems like precision of 10 is ok,
// i.e. plots looks reasonable and meaningful
func round(n float64, precision int) float64 {
	pbuf := math.Pow(10, float64(precision))
	buf := int(n * pbuf)
	return float64(buf) / pbuf
}
