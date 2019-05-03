// Package layouts handles all panel related widgets.
package layouts

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GDesktop is the exported implementation of the desktop widget
type GDesktop interface{}

/*
gdesktop widget is the representation of a form's desktop.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.

	gdesktop is the root pane, it contains components (mainly panels, these fill the available space), dialogs
	(centered), combolists, popupmenus, and tooltips. You can add only components and dialogs to desktop, the last will
	be on the top. Desktop is also suitable to create MDI applications.
*/
type gdesktop struct {
	component w.GComponent
}

// GDesktop returns a desktop object with default values.
func GDesktop() *GDesktop {
	return gdesktop{}
}
