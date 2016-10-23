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
	_ "github.com/paulrosania/go-charset/charset/iconv"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

var (
	decoder *encoding.Decoder
)

func init() {
	decoder = charmap.Windows1251.NewDecoder()
}

/* I just can't believe that someone still using cp1251.
 * Burn in hell motherfuckers! */
func StringConvert(s string) (string, error) {
	return decoder.String(s)
}
