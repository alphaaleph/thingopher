// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	"encoding/xml"

	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GButton supports push button single functionality to trigger an event.
type GButton interface {
	Action()
}

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
		Text		: The text string that the button displays.

	The following button displays a short text string and an image. Its content is on the left and centered vertically.
	The second button resembles an HTML link, and when pressed invokes the method goHome.

		<gbutton text="Button" icon="image.gif" alignment="left" tooltip="ToolTip" />
		<gbutton text="www.thinlet.com" type="link" action="goHome" />
*/
type gbutton struct {
	w.GComponent
	XMLName   xml.Name    `xml:"gbutton"`
	Icon      p.Icon      `xml:"icon,attr"`
	Mnemonic  int         `xml:"mnemonic,attr"`
	Alignment p.Alignment `xml:"alignment,attr"`
	Category  p.Category  `xml:"category,attr"`
	Text      string      `xml:"text,attr"`
}

// NewButton returns a button structure with default values.
func NewButton() GButton {
	return &gbutton{Alignment: p.ACenter, Category: p.CNormal, Mnemonic: -1}
}

// Listeners

// Action invokes the given method when the button is pressed by mouse or keyboard.
func (gb gbutton) Action() {

}
