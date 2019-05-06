// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GList displays a list of information.
type GList interface{}

/*
glist parameters:

		GComponent	: @see gcomponent.
		selection	: Possible values are: single, interval, and multiple. The default single value allows to select
						one list item at a time, the interval value to select one contiguous range of items, and the
						multiple value to select one or more contiguous ranges of items. Default = single.
		line	  	: If present and set to false, don't draw lines separating the list items. Default = true.

	A list allows the user to select one or more objects from a list, and internally handles scrolling. The following
	example list has three items, the first is selected and the last is disabled, allows to select multiple items.

		<glist selection="multiple">
 	 		<gitem text="List" selected="true" />
 	 		<gitem text="Item"/>
 	 		<gitem text="Disabled" enabled="false" />
		</glist>
*/
type glist struct {
	w.GComponent
	selection p.Selection
	line      bool
}

// NewList returns a list structure with default values.
func NewList() GList {
	return &glist{selection: p.Single, line: true}
}
