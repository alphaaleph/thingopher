// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GDialog is the exported implementation of the dialog widget
type GDialog interface{}

/*
gdialog widget is the representation of a form's dialog.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		panel		: columns, top, left, bottom, right, and gap parameters from gpanel.
		icon        : An image to be displayed in the titlebar of the dialog.
		modal       : A modal dialog grabs all the input to the components behind the dialog from the user.
						Default = false.
		resizable   : Set whether the dialog can be resized. Default = false.
		closable    : Set whether the dialog can be closed by a button. Default = false.
		maximizable : Set whether the dialog can be maximized by a button. Default = false.
		iconifiable	: Set whether the dialog can be iconified by a button. Default = false.
		text        : The title is displayed in the title bar.

	gdialog is similar to panel, but it has border and title, and you can drag it.
*/
type gdialog struct {
	component   w.GComponent
	Panel       GPanel
	icon        p.GIcon
	modal       bool
	resizable   bool
	closable    bool
	maximizable bool
	iconifiable bool
	text        string
}

// GDialog returns a dialog object with default values.
func GDialog() *GDialog {
	return gdialog{modal: false, resizable: false, closable: false, maximizable: false, iconifiable: false}
}
