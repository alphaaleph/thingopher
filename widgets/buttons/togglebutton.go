// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GToggleButton is used to display checked/unchecked states.
type GToggleButton interface{}

/*
gtogglebutton parameters:

		GComponent	: @see gcomponent.
		GCheckBox	: @see gcheckbox.

	An implementation of a two-state button, behaves as checkbox.

		<gtogglebutton text="ToggleButton" icon="image.gif" selected="true"/>
		<gtogglebutton text="ToggleButton" group="group"/>
*/
type gtogglebutton struct {
	w.GComponent
	GCheckBox
}

// NewToggleButton returns a toggle button structure with default values.
func NewToggleButton() GToggleButton {
	return &gtogglebutton{}
}
