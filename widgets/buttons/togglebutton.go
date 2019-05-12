// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	w "ebonynaranja.com/thingopher/widgets"
)

/*
GToggleButton is used to display checked/unchecked states.

	An implementation of a two-state button, behaves as checkbox.

		{"gtogglebutton":{"text":"ToggleButton","icon":"image.gif","selected":"true"}}
		{"gtogglebutton":{"text":"ToggleButton","group":"group"}}

	Parameters:
		GCheckBox	: @see GCheckBox.
		GComponent	: @see GComponent.

	Keyboard:
		Spacebar	: Selects or deselects.
		Tab			: Navigate forward.
		ShiftTab	: Navigate backward.
*/
type GToggleButton struct {
	GCheckBox
	w.GComponent
}

// DefaultGToggleButton returns a toggle button structure with default values.
func DefaultGToggleButton() *GToggleButton {
	return &GToggleButton{}
}
