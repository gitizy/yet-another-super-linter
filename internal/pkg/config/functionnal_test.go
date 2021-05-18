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

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig_GetEnabled(t *testing.T) {
	tests := []struct {
		config   Config
		expected map[string]LinterSettings
	}{
		{
			expected: map[string]LinterSettings{},
		},
		{
			config: Config{
				Linters: Linters{Disable: []string{"shellcheck"}},
			},
			expected: map[string]LinterSettings{},
		},
		{
			config: Config{
				Linters: Linters{Enable: []string{"shfmt", "shellcheck"}, Disable: []string{"shellcheck"}},
				LintersSettings: LintersSettings{
					map[string]LinterSettings{
						"shfmt": {Command: []string{"shfmt"}},
					},
				},
			},
			expected: map[string]LinterSettings{
				"shfmt": {Command: []string{"shfmt"}},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := test.config.EnabledLinters()
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("expected: %+v, but got: %+v", test.expected, got)
			}
		})
	}
}
