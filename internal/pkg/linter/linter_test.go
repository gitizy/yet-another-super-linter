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
	"fmt"
	"reflect"
	"testing"
	"testing/quick"

	"code.gitizy.dev/yasl/internal/pkg/config"
)

func Test_linter_GetName(t *testing.T) {
	f := func(name string) bool {
		l := New(name, config.LinterSettings{})
		return l.GetName() == name
	}

	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func Example_linter_PrintVersion() {
	l := New("no version", config.LinterSettings{})
	if err := l.PrintVersion(); err != nil {
		panic(err)
	}

	l = New("foo bar", config.LinterSettings{Version: []string{"echo", "foo", "bar"}})
	if err := l.PrintVersion(); err != nil {
		panic(err)
	}

	l = New("exit 1", config.LinterSettings{Version: []string{"sh", "-c", "echo 'The exit 1 version'; exit 1"}})
	if err := l.PrintVersion(); err == nil || err.Error() != "cannot print 'exit 1' version: exit status 1" {
		panic(fmt.Sprint("not expected error: ", err))
	}

	// Output:
	// ================================================================================
	// ===== Version of no version ====================================================
	// ================================================================================
	// ================================================================================
	// ===== Version of foo bar =======================================================
	// ================================================================================
	// foo bar
	// ================================================================================
	// ===== Version of exit 1 ========================================================
	// ================================================================================
	// The exit 1 version
}

func Example_linter_ExecuteLinter() {
	runLinter(New(
		"no files",
		config.LinterSettings{Command: []string{"echo"}},
	))

	runLinter(New(
		"files pattern",
		config.LinterSettings{
			Command: []string{"echo", "one"},
			Filters: config.Filters{Files: []string{"two", "three"}},
		},
	))

	runLinter(New(
		"folder pattern",
		config.LinterSettings{
			Command: []string{"echo", "one"},
			Filters: config.Filters{
				Folders: []string{"two", "three"},
			},
		},
	))

	runLinter(New(
		"git pattern",
		config.LinterSettings{
			Command: []string{"echo", "foo"},
			Filters: config.Filters{
				GitPattern: []string{"**"},
			},
		},
	))

	runLinter(New(
		"find pattern",
		config.LinterSettings{
			Command: []string{"echo", "foo"},
			Filters: config.Filters{
				FindPattern: []string{"*"},
			},
		},
	))

	runBadLinter(
		New(
			"git error",
			config.LinterSettings{
				Command: []string{"echo", "foo"},
				Filters: config.Filters{
					GitPattern: []string{""},
				},
			},
		),
		"cannot prepare linter 'git error' execution: exit status 128",
		-1,
	)

	runBadLinter(
		New(
			"bad linter",
			config.LinterSettings{
				Command: []string{"/nonexistant"},
				Filters: config.Filters{
					Folders: []string{"."},
				},
			},
		),
		"cannot execute linter 'bad linter': fork/exec /nonexistant: no such file or directory",
		-1,
	)

	{
		returnCode, err := New(
			"failed linter",
			config.LinterSettings{
				Command: []string{"sh"},
				Filters: config.Filters{
					Folders: []string{"-c", "exit 42"},
				},
			},
		).ExecuteLinter()
		if err != nil {
			panic(err)
		}
		if returnCode != 42 {
			panic(fmt.Sprint("bad return code: ", returnCode))
		}
	}

	// Output:
	// ================================================================================
	// ===== Execute no files =========================================================
	// ================================================================================
	// No file to lint. Just skip linter: no files
	// ================================================================================
	// ===== Execute files pattern ====================================================
	// ================================================================================
	// one three two
	// ================================================================================
	// ===== Execute folder pattern ===================================================
	// ================================================================================
	// one three two
	// ================================================================================
	// ===== Execute git pattern ======================================================
	// ================================================================================
	// foo linter.go linter_test.go
	// ================================================================================
	// ===== Execute find pattern =====================================================
	// ================================================================================
	// foo ./linter.go ./linter_test.go
	// ================================================================================
	// ===== Execute git error ========================================================
	// ================================================================================
	// ================================================================================
	// ===== Execute bad linter =======================================================
	// ================================================================================
	// ================================================================================
	// ===== Execute failed linter ====================================================
	// ================================================================================
}

func runLinter(l Linter) {
	if returnCode, err := l.ExecuteLinter(); err != nil {
		panic(err)
	} else if returnCode != 0 {
		panic(fmt.Sprint("Unexpected return code: ", returnCode))
	}
}

func runBadLinter(l Linter, errorText string, expectedReturnCode int) {
	returnCode, err := l.ExecuteLinter()
	if err == nil || err.Error() != errorText {
		panic(fmt.Sprint("bad error catched: ", err))
	}
	if returnCode != expectedReturnCode {
		panic(fmt.Sprint("bad return code: ", returnCode))
	}
}

func Test_zeroSplitter(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []string
	}{
		{name: "nil"},
		{name: "empty", input: []byte{}},
		{
			name:  "zero length",
			input: []byte{0},
			want:  []string{""},
		},
		{
			name:  "one",
			input: []byte("one\000"),
			want:  []string{"one"}},
		{
			name:  "one no end",
			input: []byte("one"),
			want:  []string{"one"},
		},
		{
			name:  "two",
			input: []byte("one\000two\000"),
			want:  []string{"one", "two"},
		},
		{
			name:  "two no end",
			input: []byte("one\000two"),
			want:  []string{"one", "two"},
		},
		{
			name:  "first zero length",
			input: []byte("\000two\000three\000"),
			want:  []string{"", "two", "three"},
		},
		{
			name:  "middle zero length",
			input: []byte("one\000\000three\000"),
			want:  []string{"one", "", "three"},
		},
		{
			name:  "last zero length",
			input: []byte("one\000two\000\000"),
			want:  []string{"one", "two", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroSplitter(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroSplitter() = %v, want %v", got, tt.want)
			}
		})
	}
}
