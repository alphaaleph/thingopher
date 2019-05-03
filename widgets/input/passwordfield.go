// Package input handles all widgets that are used to collect information.
package input

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GPasswordField is the exported implementation of the password field widget
type GPasswordField interface{}

/*
gpasswordfield widget is the representation of a form's data collection password field.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		textfield	: text, columns, editable, and aligment parameters from gtextfield.

	PasswordField is similar to TextField component, but does not show the original characters.

		<gpasswordfield text="secret value" />
*/
type gpasswordfield struct {
	component w.GComponent
	textfield GTextField
}

// GPasswordField returns a password field object with default values.
func GPasswordField() *GPasswordField {
	return gpasswordfield{}
}
