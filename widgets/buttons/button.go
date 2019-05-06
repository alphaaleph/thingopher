// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GButton supports push button single functionality to trigger an event.
type GButton interface{}

/*
gbutton parameters:

		GComponent	: @see gcomponent.
		icon		: The icon image that the checkbox displays.
		mnemonic 	: Specifies the index of underlined char and a key combination (Alt + the char) which invokes the
						(not necessarily focused, but enabled and visible) button's action listener. Default = -1.
		alignment	: The alignment of the text and image similar to label. Possible values are: center, left, and
						right. The default value is center. Icons are always displayed to the left of text (if any).
						Default = center.
		category	: Possible values are: normal, default, cancel, and link. The default value is normal. Default, and
						cancel values are for dialog control. Link changes the appearance of button so that it
						resembles HTML link. Default = normal.
		text		: The text string that the button displays.

	The following button displays a short text string and an image. Its content is on the left and centered vertically.
	The second button resembles an HTML link, and when pressed invokes the method goHome.

		<gbutton text="Button" icon="image.gif" alignment="left" tooltip="ToolTip" />
		<gbutton text="www.thinlet.com" type="link" action="goHome" />
*/
type gbutton struct {
	w.GComponent
	icon      p.Icon
	mnemonic  int
	alignment p.Alignment
	category  p.Category
	text      string
}

// NewButton returns a button structure with default values.
func NewButton() GButton {
	return &gbutton{alignment: p.ACenter, category: p.CNormal, mnemonic: -1}
}
