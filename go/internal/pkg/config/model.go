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

package config

// Config is the root of a configuration.
type Config struct {
	Linters         Linters
	LintersSettings LintersSettings `mapstructure:"linters-settings"`
}

// Linters say which linter is enabled or disabled.
type Linters struct {
	Enable  []string
	Disable []string
}

// LintersSettings is just a wrapper for all the linters settings.
type LintersSettings struct {
	Settings map[string]LinterSettings `mapstructure:",remain"`
}

// LinterSettings hold one linter configuration.
type LinterSettings struct {
	Version []string
	Command []string
	Args    []string
	Workdir string
	Filters Filters
}

// Filters limit the linted files with various filtering patterns.
// Some linter may not use all patterns.
type Filters struct {
	GitPattern  []string `mapstructure:"git-pattern"`
	FindPattern []string `mapstructure:"find-pattern"`
	Files       []string
	Folders     []string
}
