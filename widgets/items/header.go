// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

// GHeader holds the information for the columns in a table.
type GHeader interface{}

/*
gheader parameters:

		columns	: @see gcolumn.

	Table header includes the columns. They define the column texts and widths.

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
type gheader struct {
	columns []GColumn
}

// NewHeader returns a header structure with default values.
func NewHeader() GHeader {
	return &gheader{}
}
