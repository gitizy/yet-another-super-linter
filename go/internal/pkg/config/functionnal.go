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

// EnabledLinters return the list of effectively enabled linters.
func (c Config) EnabledLinters() map[string]LinterSettings {
	result := make(map[string]LinterSettings)
	for _, linter := range c.Linters.listEnabled() {
		result[linter] = c.LintersSettings.Settings[linter]
	}

	return result
}

func (l Linters) listEnabled() []string {
	status := make(map[string]bool)

	for _, v := range l.Enable {
		status[v] = true
	}

	for _, v := range l.Disable {
		status[v] = false
	}

	var res []string

	for name, isActive := range status {
		if isActive {
			res = append(res, name)
		}
	}

	return res
}
