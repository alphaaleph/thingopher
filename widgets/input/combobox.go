// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GComboBox is the exported implementation of the combo box widget
type GComboBox interface{}

/*
gcombobox widget is the representation of a form's data collection combo box.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		textfield	: text, columns, editable, and aligment parameters from gtextfield.
		icon		: he icon image that the combobox displays.
		selected	: The index of the currently selected choice, value -1 indicates a custom edited value in an
						editable box. Default = -1.

	gcombobox is a combination of a text field and drop-down list. This example has a default value, two choices, and
	it is editable.

		<gcombobox text="ComboBox">
 	 		<choice text="Choice" icon="image.gif" />
 	 		<choice text="Disabled" enabled="false" />
		</gcombobox>
*/
type gcombobox struct {
	component w.GComponent
	textfield GTextField
	icon      p.GIcon
	selected  int
}

// GComboBox returns a combo box object with default values.
func GComboBox() *GComboBox {
	return gcombobox{selected: -1}
}
