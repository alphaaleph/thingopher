// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GDialog is similar to panel, but it has a border and title, and you can drag it.
type GDialog interface{}

/*
gdialog parameters:

		GComponent	: @see gcomponent.
		GPanel		: @see gpanel.
		icon        : An image to be displayed in the titlebar of the dialog.
		modal       : A modal dialog grabs all the input to the components behind the dialog from the user.
						Default = false.
		resizable   : Set whether the dialog can be resized. Default = false.
		closable    : Set whether the dialog can be closed by a button. Default = false.
		maximizable : Set whether the dialog can be maximized by a button. Default = false.
		iconifiable	: Set whether the dialog can be iconified by a button. Default = false.
		text        : The title is displayed in the title bar.
*/
type gdialog struct {
	w.GComponent
	GPanel
	icon        p.Icon
	modal       bool
	resizable   bool
	closable    bool
	maximizable bool
	iconifiable bool
	text        string
}

// NewDialog returns a dialog structure with default values.
func NewDialog() GDialog {
	return &gdialog{modal: false, resizable: false, closable: false, maximizable: false, iconifiable: false}
}
