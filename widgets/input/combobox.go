// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GComboBox supports a text field that has a drop-down list.
type GComboBox interface{}

/*
gcombobox parameters:

		GComponent	: @see gcomponent.
		GTextField	: @see gtextfield.
		icon		: he icon image that the combobox displays.
		selected	: The index of the currently selected choice, value -1 indicates a custom edited value in an
						editable box. Default = -1.

	gcombobox is a combination of a text field and drop-down list. This example has a default value, two choices, and
	it is editable.

		<gcombobox text="ComboBox">
 	 		<gchoice text="Choice" icon="image.gif" />
 	 		<gchoice text="Disabled" enabled="false" />
		</gcombobox>
*/
type gcombobox struct {
	w.GComponent
	GTextField
	icon     p.Icon
	selected int
}

// NewComboBox returns a combo box structure with default values.
func NewComboBox() GComboBox {
	return &gcombobox{selected: -1}
}
