// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GMenuBar offers a selection of choices based on different contexts.
type GMenuBar interface{}

/*
gmenubar parameters:

		GComponent	: @see gcomponent.
		placement	: Menubar may unfold the menus either downwards (default - for menubars placed at the top of the
						container), or upwards from the menubar position (bottom - for menubars placed at the bottom of
						the container).

	The following panel contains a menubar on the top and a large text-area. The menubar has only one menu ('Menu'), it
	has a menuitem ('MenuItem') and a menu ('Menu'). The later's popup contains a menuitem with a selected checkbox
	('CheckBox'), a separator, and a radio button ('RadioButton').

		<gpanel columns="1" gap="4">
  			<gmenubar weightx="1">
    			<gmenu text="Menu">
      				<gmenuitem text="MenuItem" icon="image.gif" />
      				<gmenu text="Menu">
        				<gcheckboxmenuitem text="CheckBox" selected="true" />
        				<gseparator />
        				<gcheckboxmenuitem text="RadioButton" group="group"/>
      				</gmenu>
    			</gmenu>
  			</gmenubar>
  			<gtextarea weighty="1" />
		</gpanel>
*/
type gmenubar struct {
	w.GComponent
	placement p.Placement
}

// NewMenuBar returns a menu bar structure with default values.
func NewMenuBar() GMenuBar {
	return &gmenubar{placement: p.Top}
}
