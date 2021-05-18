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

package utils

func ExampleTitle() {
	PrintTitle("My simple title")
	PrintTitle("My just long enough title. The longer is exactly 80 characters long !!!!!!")
	PrintTitle("My very very very very very very long title, this is much longer than 80 characters !!!")

	// Output:
	// ================================================================================
	// ===== My simple title ==========================================================
	// ================================================================================
	// ================================================================================
	// ===== My just long enough title. The longer is exactly 80 characters long !!!!!!
	// ================================================================================
	// ================================================================================
	// ===== My very very very very very very long title, this is much longer than 80 characters !!!
	// ================================================================================
}
