// Package buttons manage all related button widgets
package buttons

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GToggleButton is the exported implementation of the toggle button widget
type GToggleButton interface{}

/*
gtogglebutton widget is the representation of a form's toggle button.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from component.
		GCheckBox	: see GCheckBox parameters.

	An implementation of a two-state button, behaves as checkbox.

		<gtogglebutton text="ToggleButton" icon="image.gif" selected="true"/>
		<gtogglebutton text="ToggleButton" group="group"/>
*/
type gtogglebutton struct {
	component w.GComponent
	GCheckBox
}

// GToggleButton returns a toggle button object with default values.
func GToggleButton() *GToggleButton {
	return gtogglebutton{selected: false, alignment: p.ACenter, mnemonic: -1}
}
