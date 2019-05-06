// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

// GRow is an instance of a single row in a table.
type GRow interface{}

/*
grow parameters:

		selected	: True if currently selected the entire row. Default = false.
		cells		: @see gcell.

	Table contains rows (and columns for the header), and row contains cells. The 'row' key identifies the list of
	rows. An entire row (not cell) is selectable in a table.

		<gtable selection="multiple">
 	 		<gheader>
 		   		<gcolumn text="A" width="24" /><gcolumn text="B" />
 	 		</gheader>
 	 		<grow selected="true">
 		   		<gcell text="a1" /><gcell text="b1" enabled="false" />
 	 		</grow>
 	 		<grow>
 		   		<gcell text="a2" /><gcell text="b2" />
 	 		</grow>
		</gtable>
*/
type grow struct {
	selected bool
	cells    []GCell
}

// NewRow returns a row structure with default values.
func NewRow() GRow {
	return &grow{selected: false}
}
