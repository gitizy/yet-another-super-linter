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

package linter

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"code.gitizy.dev/yasl/internal/pkg/config"
	"code.gitizy.dev/yasl/internal/pkg/utils"
)

// Linter represent an executable linter.
type Linter interface {
	GetName() string
	PrintVersion() error
	ExecuteLinter() (returnCode int, err error)
}

type linter struct {
	name   string
	config config.LinterSettings
}

// New instantiate an linter with the given configuration.
func New(name string, cfg config.LinterSettings) Linter {
	return &linter{
		name:   name,
		config: cfg,
	}
}

func (l *linter) GetName() string {
	return l.name
}

func (l *linter) PrintVersion() error {
	utils.PrintTitle(fmt.Sprintf("Version of %s", l.name))
	if len(l.config.Version) == 0 {
		return nil
	}

	var args []string
	if len(l.config.Version) > 1 {
		args = append(args, l.config.Version[1:]...)
	}
	// #nosec G204
	cmd := exec.Command(l.config.Version[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("cannot print '%s' version: %v", l.name, err)
	}

	return nil
}

func (l *linter) ExecuteLinter() (int, error) {
	utils.PrintTitle(fmt.Sprintf("Execute %s", l.name))

	cmd, err := l.buildLinterCommand()
	if err != nil {
		return -1, fmt.Errorf("cannot prepare linter '%s' execution: %v", l.name, err)
	}
	if cmd == nil {
		fmt.Println("No file to lint. Just skip linter:", l.name)
		return 0, nil
	}

	err = cmd.Run()

	if err != nil {
		exitCode := &exec.ExitError{}
		if errors.As(err, &exitCode) {
			return exitCode.ExitCode(), nil
		}
		return -1, fmt.Errorf("cannot execute linter '%s': %v", l.name, err)
	}

	return 0, nil
}

func (l *linter) buildLinterCommand() (*exec.Cmd, error) {
	var args []string
	if len(l.config.Command) > 1 {
		args = append(args, l.config.Command[1:]...)
	}
	args = append(args, l.config.Args...)

	files, err := l.listFiles()
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, nil
	}
	args = append(args, files...)

	// #nosec G204
	cmd := exec.Command(l.config.Command[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = l.config.Workdir

	return cmd, nil
}

func (l *linter) listFiles() ([]string, error) {
	var result []string

	result = append(result, l.config.Filters.Files...)
	result = append(result, l.config.Filters.Folders...)

	gitFiles, err := l.listGitFiles()
	if err != nil {
		return nil, err
	}
	result = append(result, gitFiles...)

	findFiles, err := l.listFindFiles()
	if err != nil {
		return nil, err
	}
	result = append(result, findFiles...)

	return result, nil
}

func (l *linter) listGitFiles() ([]string, error) {
	if len(l.config.Filters.GitPattern) == 0 {
		return nil, nil
	}

	args := []string{"ls-files", "-z", "--"}
	args = append(args, l.config.Filters.GitPattern...)

	cmd := exec.Command("git", args...)
	raw, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return zeroSplitter(raw), nil
}

func (l *linter) listFindFiles() ([]string, error) {
	if len(l.config.Filters.FindPattern) == 0 {
		return nil, nil
	}

	args := []string{".", "-type", "f"}
	for _, pattern := range l.config.Filters.FindPattern {
		args = append(args, "-name", pattern)
	}
	args = append(args, "-print0")

	cmd := exec.Command("find", args...)
	raw, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return zeroSplitter(raw), nil
}

func zeroSplitter(raw []byte) []string {
	if len(raw) == 0 {
		return nil
	}

	var result []string

	if raw[len(raw)-1] != 0 {
		raw = append(raw, 0)
	}

	for len(raw) > 0 {
		for i, b := range raw {
			if b == 0 {
				result = append(result, string(raw[:i]))
				raw = raw[i+1:]
				break
			}
		}
	}

	return result
}
