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
	"errors"
	"reflect"
	"testing"
)

func TestConfig_UnmarshalConfig(t *testing.T) {
	data := []byte(`
---
linters:
  enable: [shellcheck, shfmt]
linters-settings:
  golangci-lint:
    version: ['golangci-lint', '--version']
    command: ['golangci-lint']
    args:
      - '--config=/opt/yasl/linters/golangci-lint.yaml'
    filters:
      folders:
        - .
  shellcheck:
    version: ['shellcheck', '--version']
    command: ['shellcheck']
    args:
      - '--external-sources'
      - '--enable=quote-safe-variables'
      - '--enable=require-variable-braces'
    filters:
      git-pattern: ['**.sh']
`)

	expected := Config{
		Linters: Linters{
			Enable: []string{"shellcheck", "shfmt"},
		},
		LintersSettings: LintersSettings{
			Settings: map[string]LinterSettings{
				"golangci-lint": {
					Version: []string{"golangci-lint", "--version"},
					Command: []string{"golangci-lint"},
					Args:    []string{"--config=/opt/yasl/linters/golangci-lint.yaml"},
					Filters: Filters{
						Folders: []string{"."},
					},
				},
				"shellcheck": {
					Version: []string{"shellcheck", "--version"},
					Command: []string{"shellcheck"},
					Args: []string{
						"--external-sources",
						"--enable=quote-safe-variables",
						"--enable=require-variable-braces",
					},
					Filters: Filters{
						GitPattern: []string{"**.sh"},
					},
				},
			},
		},
	}

	var c Config
	if err := Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("expected: %#v, but got: %#v", expected, c)
	}

	invalidYamlError := &InvalidYamlError{}
	if err := Unmarshal([]byte("`"), nil); !errors.As(err, &invalidYamlError) {
		t.Fatalf("Error not catched on bad yaml: %v", err)
	} else {
		t.Logf("Catched bad yaml error: %v", err)
	}

	invalidStructureError := &InvalidStructureError{}
	if err := Unmarshal([]byte("linters: 3"), nil); !errors.As(err, &invalidStructureError) {
		t.Fatalf("Error not catched on bad structure: %v", err)
	} else {
		t.Logf("Catched bad structure error: %v", err)
	}
}
