// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import (
	p "ebonynaranja.com/thingopher/properties"
	d "ebonynaranja.com/thingopher/widgets/display"
)

// GColumn is an instance of a single column in a table.
type GColumn interface{}

/*
gcolumn parameters:

		GLabel	: @see glabel.
		width	: The width of the column. The last (rightmost) column always extends to fill in the remaining width of
					the table, as determined by the parent container. Default = 80.
		sort	: If present, and set to either "ascent" or "descent", an arrow is drawn near right end of the column,
					as a visual indication of sorting order. NOTE: for now this is just a visual decoration,
					application still needs to do the sorting. Default = none.

	Table header includes the columns. They define the column texts and widths. If the total sum of column widths
	exceeds table bounds set by a parent container, horizontal scrollbar is drawn.

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
type gcolumn struct {
	d.GLabel
	width int
	sort  p.Sort
}

// NewColumn returns a column structure with default values.
func NewColumn() GColumn {
	return &gcolumn{width: 80, sort: p.None}
}
