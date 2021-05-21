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

package lint

import (
	"fmt"
	"io/ioutil"
	"sort"

	"code.gitizy.dev/yasl/internal/pkg/config"
	"code.gitizy.dev/yasl/internal/pkg/linter"
	"code.gitizy.dev/yasl/internal/pkg/utils"
)

// Entrypoint execute the lint command
func Entrypoint(cfgFile string) (succeed bool) {
	linters, err := loadLinters(cfgFile)
	if err != nil {
		panic(err)
	}

	return runLinters(linters)
}

func loadLinters(cfgFile string) ([]linter.Linter, error) {
	raw, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	var cfg config.Config
	if err = config.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}

	var names []string
	for name := range cfg.EnabledLinters() {
		names = append(names, name)
	}
	sort.Strings(names)

	enabled := cfg.EnabledLinters()
	var linters []linter.Linter
	for _, name := range names {
		linters = append(linters, linter.New(name, enabled[name]))
	}

	return linters, nil
}

func runLinters(linters []linter.Linter) (succeed bool) {
	var failed []string

	for _, linter := range linters {
		_ = linter.PrintVersion()

		returnCode, err := linter.ExecuteLinter()
		if err != nil {
			fmt.Println("Execution failed: ", err)
		} else if returnCode > 0 {
			fmt.Println("Execution failed with return code: ", returnCode)
		}

		if returnCode != 0 {
			failed = append(failed, linter.GetName())
		}
	}

	if len(failed) > 0 {
		utils.PrintTitle("failed linters")

		for _, name := range failed {
			fmt.Println(" - ", name)
		}
	}

	return len(failed) == 0
}
