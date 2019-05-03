// Package input handles all widgets that are used to collect information.
package input

import (
	"math"

	w "ebonynaranja.com/thingopher/widgets"
)

// GSpinBox is the exported implementation of the spin box widget
type GSpinBox interface{}

/*
gspinbox widget is the representation of a form's data collection spin box.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		textfield	: text, columns, editable, and aligment parameters from gtextfield.
		minimum		: Minimum value. It's not possible to set the spinbox below this value by the arrow button.
		maximum		: Maximum value. It's not possible to set the spinbox above this value by the arrow button.
		step		: Increment step when using spinbox arrows. Default = 1.

	gspinbox is a textfield width 2 arrow buttons to change (increase, and decrease by 1) its integer value. The
	textfield input is filtered for digits, but the range isn't checked, you can also delete its content.

		<gspinbox text="100"/>
*/
type gspinbox struct {
	component w.GComponent
	textfield GTextField
	minimum   int
	maximum   int
	step      int
}

// GSpinBox returns a spin box object with default values.
func GSpinBox() *GSpinBox {
	return gspinbox{minimum: math.MinInt32, maximum: math.MaxInt32, step: 1}
}
