// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	"fmt"

	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

/*
GButton supports push button single functionality to trigger an event.

	The following button displays a short text string and an image. Its content is on the left and centered vertically.
	The second button resembles an HTML link, and when pressed invokes the method goHome.

		{"gbutton":{"text":"Button","icon":"image.gif","alignment":"left","tooltip":"ToolTip"}}
		{"gbutton":{"text":"www.thinlet.com","type":"link","action":"goHome"}}

	Properties:
		Icon		: The icon image that the checkbox displays.
		Mnemonic 	: Specifies the index of underlined char and a key combination (Alt + the char) which invokes the
						(not necessarily focused, but enabled and visible) button's action listener. Default = -1.
		Alignment	: The alignment of the text and image similar to label. Possible values are: center, left, and
						right. The default value is center. Icons are always displayed to the left of text (if any).
						Default = center.
		Type		: Possible values are: normal, default, cancel, and link. The default value is normal. Default, and
						cancel values are for dialog control. Link changes the appearance of button so that it
						resembles HTML link. Default = normal.
		Text		: The text string that the button displays.
		GComponent	: @see GComponent.

	Listeners:
		Action		: Invokes the given method when the button is pressed by mouse or keyboard.

	Keyboard:
		Spacebar	: Activates button.
		Tab			: Navigate forward.
		ShiftTab	: Navigate backward.
*/
type GButton struct {
	Icon      p.Icon      `json:"icon"`
	Mnemonic  int         `json:"mnemonic"`
	Alignment p.Alignment `json:"alignment"`
	Type      p.Type      `json:"type"`
	Text      string      `json:"text"`
	w.GComponent
}

// DefaultGButton returns a button structure with default values.
func DefaultGButton() *GButton {
	return &GButton{Alignment: p.ACenter, Type: p.TNormal, Mnemonic: -1}
}

// Action invokes the given method when the button is pressed by mouse or keyboard.
func (gb GButton) Action() {
	fmt.Println(gb)
}
