// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTextField enables the user to enter text information.
type GTextField interface{}

/*
gtextfield parameters:

		GComponent	: @see gcomponent
		end			: End index of the selection, same as the caret position. Default = 0.
		start		: Start index of the selection. Default = 0.
		columns		: The preferred width of the component is fixed (if 0) or calculated by the given value.
						Default = 0.
		alignment	: The alignment of the text content along the X axis. Possible values are: left, center, and right.
						Default = left.
		editable	: The specified boolean to indicate whether or not this textfield should be editable. A
						non-editable field is focusable, and selectable. Default = true.
		text		: The text contained in this textarea. Default = ""

	gtextfield component allows the editing of a single line of text. This field has a simple text value and 10
	characters width.

		<gtextfield text="TextField" columns="10" />
*/
type gtextfield struct {
	w.GComponent
	end       uint
	start     uint
	columns   uint16
	alignment p.Alignment
	editable  bool
	text      string
}

// NewTextField returns a text field structure with default values.
func NewTextField() GTextField {
	return &gtextfield{alignment: p.ALeft, editable: true}
}
