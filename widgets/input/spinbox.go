// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	"math"

	w "ebonynaranja.com/thingopher/widgets"
)

// GSpinBox is a text box that accepts a range of values.
type GSpinBox interface{}

/*
gspinbox parameters:

		GComponent	: @see gcomponent.
		GTextField	: @see gtextfield.
		minimum		: Minimum value. It's not possible to set the spinbox below this value by the arrow button.
		maximum		: Maximum value. It's not possible to set the spinbox above this value by the arrow button.
		step		: Increment step when using spinbox arrows. Default = 1.

	gspinbox is a textfield width 2 arrow buttons to change (increase, and decrease by 1) its integer value. The
	textfield input is filtered for digits, but the range isn't checked, you can also delete its content.

		<gspinbox text="100"/>
*/
type gspinbox struct {
	w.GComponent
	GTextField
	minimum int
	maximum int
	step    int
}

// NewSpinBox returns a spin box structure with default values.
func NewSpinBox() GSpinBox {
	return &gspinbox{minimum: math.MinInt32, maximum: math.MaxInt32, step: 1}
}
