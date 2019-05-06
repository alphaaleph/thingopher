// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	"image/color"

	p "ebonynaranja.com/thingopher/properties"
	"golang.org/x/image/font"
)

// GChoice supports the entries listed in the list of the combo box.
type GChoice interface{}

/*
gchoice parameters:

		property	: Adds an arbitrary key/value client property (or properties) stored in a hashtable of the item.
		icon		: The icon image that the choice displays.
		font		: Custom font to use for text.
		background	: Custom background color.
		foreground	: Custom foreground (text) color.
		alignment	: The alignment of the text and image similar to label. Possible values are: left, center, and
						right.
		enabled		: Enables or disables the item. A disabled item is painted gray, and can't be selected by mouse or
						keyboard. Default = true.
		name		: Identifies the item.
		text		: The text text that the choice displays.
		tooltip		: The text pops up when the mouse lingers inside this specified part of the component, otherwise
						the tooltip text of the component.

	The drop down list of a combobox contains choices. The popup combolist internally handles scrolling.

		<gcombobox text="ComboBox">
 	 		<gchoice text="Choice" icon="image.gif" />
 	 		<gchoice text="Disabled" enabled="false" />
		</gcombobox>
*/
type gchoice struct {
	p.Property
	icon       p.Icon
	font       font.Style
	background color.Color
	foreground color.Color
	alignment  p.Alignment
	enabled    bool
	name       string
	text       string
	tooltip    string
}

// NewChoice returns a choice structure with default values.
func NewChoice() GChoice {
	return &gchoice{alignment: p.ALeft, enabled: true}
}
