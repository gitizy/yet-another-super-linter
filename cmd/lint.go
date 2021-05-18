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
	"os"

	"code.gitizy.dev/yasl/internal/app/lint"
	"github.com/spf13/cobra"
)

var cfgFile string

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Run linters",
	Run: func(cmd *cobra.Command, args []string) {
		if !lint.Entrypoint(cfgFile) {
			os.Exit(1)
		}
	},
}

func init() {
	lintCmd.Flags().StringVar(&cfgFile, "config", "/opt/yasl/config.yaml", "config file")
}
