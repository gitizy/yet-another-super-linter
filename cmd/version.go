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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appVersion = "unknow"
var appDate = "unknow"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display build version and date",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version:", appVersion)
		fmt.Println("date:", appDate)
		fmt.Println("license: GNU General Public License, version 3")
		fmt.Println("website: https://github.com/gitizy/yet-another-super-linter")
	},
}
