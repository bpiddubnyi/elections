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
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"io/ioutil"
	"strings"
)

/* I just can't believe that someone still using cp1251.
 * Burn in hell motherfuckers! */
func StringConvert(cp string, s string) (res string, err error) {
	r, err := charset.NewReader(cp, strings.NewReader(s))
	if err != nil {
		return
	}

	rb, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	return string(rb), nil
}
