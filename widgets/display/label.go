// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package display manage all non-input widgets.
package display

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

/*
GLabel is used to display read-only information.

	The following label displays a short text string and an image. Its content is centered both vertically and
	horizontally.

		{"glabel":{"text":"Label","icon":"image.gif","alignment":"center"}}

	Properties:
		For			: Set the component this is labelling. The label will focus the component specified by its name
						when the mnemonic is activated.
		Icon		: The icon image that the label displays.
		Mnemonic	: Specifies the underlined char in the label's text. Default = -1.
		Alignment	: The alignment of the label's content along the X axis, vertically is centered. Possible values
						are: left, center, and right. The default value, if not set, is left. The image position
						horizontally left, above, and right relative to the text. Default = left.
		Text		: The text string that the label displays.
		GComponent	: @see GComponent.
*/
type GLabel struct {
	For       *w.GComponent `json:"for"`
	Icon      p.Icon        `json:"icon"`
	Mnemonic  int           `json:"mnemonic"`
	Alignment p.Alignment   `json:"alignment"`
	Text      string        `json:"text"`
	w.GComponent
}

// DefaultGLabel returns a label structure with default values.
func DefaultGLabel() *GLabel {
	return &GLabel{Alignment: p.ALeft, Mnemonic: -1}
}
