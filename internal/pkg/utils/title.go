// Copyright (C) 2021 VERDOÏA Laurent
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package utils

import (
	"fmt"
	"strings"
)

const lineLength = 80

var padding = strings.Repeat("=", lineLength)

// PrintTitle is designed to print a beautifull 80-columned title on 3 lines.
func PrintTitle(title string) {
	line := fmt.Sprintf("===== %s", title)
	if len(line) < lineLength {
		line += " " + padding
		line = line[:80]
	}

	fmt.Println(padding)
	fmt.Println(line)
	fmt.Println(padding)
}
