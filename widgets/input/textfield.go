// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTextField is the exported implementation of the text area widget
type GTextField interface{}

/*
gtextfield widget is the representation of a form's data collection text area.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		alignment	: The alignment of the text content along the X axis. Possible values are: left, center, and right.
						Default = left.
		end			: End index of the selection, same as the caret position. Default = 0.
		start		: Start index of the selection. Default = 0.
		columns		: The preferred width of the component is fixed (if 0) or calculated by the given value.
						Default = 0.
		editable	: The specified boolean to indicate whether or not this textfield should be editable. A
						non-editable field is focusable, and selectable. Default = true.
		text		: The text contained in this textarea. Default = ""

	gtextfield component allows the editing of a single line of text. This field has a simple text value and 10
	characters width.

		<gtextfield text="TextField" columns="10" />
*/
type gtextfield struct {
	component w.GComponent
	alignment p.Alignment
	end       uint
	start     uint
	columns   uint16
	editable  bool
	text      string
}

// GTextField returns a slider object with default values.
func GTextField() *GTextField {
	return gtextfield{alignment: p.Left, editable: true}
}
