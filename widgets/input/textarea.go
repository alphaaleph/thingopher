// Package input handles all widgets that are used to collect information.
package input

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GTextArea is the exported implementation of the text area widget
type GTextArea interface{}

/*
gtextarea widget is the representation of a form's data collection text area.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		end			: End index of the selection, same as the caret position. Default = 0.
		start		: Start index of the selection. Default = 0.
		columns		: The number of visible letters in a column. Default = 0.
		rows		: The number of visible rows for this textarea. Default = 0.
		border		: If set to false, no border will be drawn around the textarea, and its background will be set to
						the default background color, to visually "blend" with its container. You can still override
						this by setting a custom background color. Default = true.
		editable	: The specified boolean to indicate whether or not this textarea should be editable. A non-editable
						area is focusable, and selectable. Default = true.
		wrap		: If set to true the lines will be wrapped at word boundaries (whitespace) if they are too long to
						fit within the allocated width. Default = false.
		text		: The text contained in this textarea. Default = ""

	TextArea is a multi-line area that displays plain text, optionally wrapped at word boundaries (whitespace), and
	internally handles scrolling.

		<gtextarea text="TextArea" wrap="true" columns="40" rows="2" />
*/
type gtextarea struct {
	component w.GComponent
	end       uint
	start     uint
	columns   uint16
	rows      uint16
	border    bool
	editable  bool
	wrap      bool
	text      string
}

// GTextArea returns a slider object with default values.
func GTextArea() *GTextArea {
	return gtextarea{border: true, editable: true, wrap: false}
}
