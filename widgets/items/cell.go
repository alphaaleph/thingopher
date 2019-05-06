// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import i "ebonynaranja.com/thingopher/widgets/input"

// GCell is a single cell in a table row.
type GCell interface{}

/*
gcell parameters:

		GComboBox	: @see gcombobox.

	Cell displays a short text and icon, its height is equals with the row's height, and its width is defined in the
	header column.

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
type gcell struct {
	i.GComboBox
}

// NewCell returns a cell structure with default values.
func NewCell() GCell {
	return &gcell{}
}
