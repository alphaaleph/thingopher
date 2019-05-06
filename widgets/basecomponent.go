// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package widgets holds all the components supported by thingopher
package widgets

import (
	p "ebonynaranja.com/thingopher/properties"
)

// GBaseComponent is the base component of all the widgets.
type GBaseComponent interface{}

/*
gbasecomponent parameters:

		property	: Binds an arbitrary key/value client property (or properties).
		enabled 	: Enables or disables this component. A disabled component is painted gray, and can't respond to
						user mouse or keyboard input, gain the keyboard focus, and generate events. Default = true.
		visble		: An invisible component doesn't take place in parent's layout. Default = true.
		name		: Identifies the component.
		tooltip		: The text pops up when the mouse lingers inside the component.
*/
type gbasecomponent struct {
	property p.Property
	enabled  bool
	visible  bool
	name     string
	tooltip  string
}

// NewBaseComponent returns a base component structure with default values.
func NewBaseComponent() GBaseComponent {
	return &gbasecomponent{enabled: true, visible: true}
}
