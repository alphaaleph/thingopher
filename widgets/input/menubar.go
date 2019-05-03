// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GMenuBar is the exported implementation of the menu bar widget
type GMenuBar interface{}

/*
gmenubar widget is the representation of a form's data collection menu bar.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		placement	: Menubar may unfold the menus either downwards (default - for menubars placed at the top of the
						container), or upwards from the menubar position (bottom - for menubars placed at the bottom of
						the container).

	The following panel contains a menubar on the top and a large text-area. The menubar has only one menu ('Menu'), it
	has a menuitem ('MenuItem') and a menu ('Menu'). The later's popup contains a menuitem with a selected checkbox
	('CheckBox'), a separator, and a radio button ('RadioButton').

		<panel columns="1" gap="4">
  			<menubar weightx="1">
    			<menu text="Menu">
      				<menuitem text="MenuItem" icon="image.gif" />
      				<menu text="Menu">
        				<checkboxmenuitem text="CheckBox" selected="true" />
        				<separator />
        				<checkboxmenuitem text="RadioButton" group="group"/>
      				</menu>
    			</menu>
  			</menubar>
  			<textarea weighty="1" />
		</panel>
*/
type gmenubar struct {
	component w.GComponent
	placement p.Placement
}

// GMenuBar returns a menu bar object with default values.
func GMenuBar() *GMenuBar {
	return gmenubar{placement: p.Top}
}
