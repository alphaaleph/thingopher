// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import i "ebonynaranja.com/thingopher/widgets/input"

// GItem is a single element in a list widget.
type GItem interface{}

/*
gitem parameters:

		GComboBox	: @see gcombobox.
		selected 	: True if the item currently selected. Default = true.

	List contains items, items display a short text string and an image and can be selected.

		<glist selection="multiple">
 	 		<gitem text="List" selected="true" />
 	 		<gitem text="Item"/>
 	 		<gitem text="Disabled" enabled="false" />
		</glist>
*/
type gitem struct {
	i.GComboBox
	selected bool
}

// NewItem returns an item structure with default values.
func NewItem() GItem {
	return &gitem{selected: true}
}
