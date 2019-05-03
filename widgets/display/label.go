// Package display manage all non-input widgets.
package display

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GLabel is the exported implementation of the label widget
type GLabel interface{}

/*
glabel widget is the representation of a form's label.

	Parameters:
		owner		: Set the component this is labelling. The label will focus the component specified by its name
						when the mnemonic is activated.
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from component.
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
	owner     *GComponent
	component w.GComponent
	icon      p.GIcon
	mnemonic  int
	alignment p.GAlignment
	text      string
}

// GLabel returns a label object with default values.
func GLabel() *GLabel {
	return glabel{alignment: p.ALeft, mnemonic: -1}
}
