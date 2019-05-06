// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GDesktop is the root pane, it contains components (mainly panels, these fill the available space), dialogs
// (centered), combolists, popupmenus, and tooltips. You can add only components and dialogs to desktop, the last will
// be on the top. Desktop is also suitable to create MDI applications.
type GDesktop interface{}

/*
gdesktop parameters:

		GComponent	: @see gcomponent.
*/
type gdesktop struct {
	w.GComponent
}

// NewDesktop returns a desktop structure with default values.
func NewDesktop() GDesktop {
	return &gdesktop{}
}
