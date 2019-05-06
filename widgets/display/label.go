// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package display manage all non-input widgets.
package display

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GLabel is used to display read-only information.
type GLabel interface{}

/*
glabel parameters:

		GComponent	: @see gcomponent.
		owner		: Set the component this is labelling. The label will focus the component specified by its name
						when the mnemonic is activated.
		icon		: The icon image that the label displays.
		mnemonic	: Specifies the underlined char in the label's text. Default = -1.
		alignment	: The alignment of the label's content along the X axis, vertically is centered. Possible values
						are: left, center, and right. The default value, if not set, is left. The image position
						horizontally left, above, and right relative to the text. Default = left.
		text		: The text string that the label displays.

	The following label displays a short text string and an image. Its content is centered both vertically and
	horizontally.

		<label text="Label" icon="image.gif" alignment="center" />
*/
type glabel struct {
	w.GComponent
	owner     *w.GComponent
	icon      p.Icon
	mnemonic  int
	alignment p.Alignment
	text      string
}

// NewLabel returns a label structure with default values.
func NewLabel() GLabel {
	return &glabel{alignment: p.ALeft, mnemonic: -1}
}
