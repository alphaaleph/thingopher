// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTable is a widget that manages data like a spreadsheet.
type GTable interface{}

/*
gtable parameters:

		GComponent	: @see gcomponent.
		GList		: @see glist.
		selection	: Possible values are: single, interval, and multiple. The default single value allows to select
						one row at a time, the interval value to select one contiguous range of rows, and the multiple
						value to select one or more contiguous ranges of rows. Default = single.

	A table presents data in a two-dimensional table format, allows to select rows, has header, and internally handles
	scrolling. The following example table includes two columns ('A', and 'B'), and two rows, the first row is
	selected, but allows to select multiple rows.

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
type gtable struct {
	w.GComponent
	GList
	selection p.Selection
}

// NewTable returns a table structure with default values.
func NewTable() GTable {
	return &gtable{selection: p.Single}
}
