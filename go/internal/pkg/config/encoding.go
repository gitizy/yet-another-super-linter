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

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// An InvalidYamlError describes a not well YAML formated invalid input passed to Unmarshal.
type InvalidYamlError struct {
	Err error
}

func (e *InvalidYamlError) Error() string {
	return fmt.Sprintf("config: %s", e.Err)
}

// An InvalidStructureError describes a not well structured config file passed to Unmarshal.
type InvalidStructureError struct {
	Err error
}

func (e *InvalidStructureError) Error() string {
	return fmt.Sprintf("config: %s", e.Err)
}

// Unmarshal parse the YAML-encoded configuration file from data and stores the result in the value pointed by config.
// If
func Unmarshal(data []byte, config *Config) error {
	raw := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return &InvalidYamlError{err}
	}

	if err := mapstructure.Decode(raw, config); err != nil {
		return &InvalidStructureError{err}
	}

	return nil
}
