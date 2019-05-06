// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GPasswordField allows the user to enter information without showing the actual characters.
type GPasswordField interface{}

/*
gpasswordfield parameters:

		GComponent	: @see gcomponent
		GTextField	: @see gtextfield

	PasswordField is similar to TextField component, but does not show the original characters.

		<gpasswordfield text="secret value" />
*/
type gpasswordfield struct {
	w.GComponent
	GTextField
}

// NewPasswordField returns a password field structure with default values.
func NewPasswordField() GPasswordField {
	return &gpasswordfield{}
}
