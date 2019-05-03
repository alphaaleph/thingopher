// Package input handles all widgets that are used to collect information.
package input

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GMenu is the exported implementation of the menu widget
type GMenu interface{}

/*
gmenu widget is the representation of a form's data collection menu.

	Parameters:
		combobox	: name, enabled, text, icon, alignment, tooltip, and property parameters similar to gcombobox.
		mnemonic	: Specifies the underlined char and the key combination which selects the tab. Default = -1.

	Menubar includes menu, and menu could contain menuitems, check or radio button items, separators, and menus.

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
type gmenu struct {
	combobox    i.GComboBox
	mnemonic    int
}

// GMenu returns a menu object with default values.
func GMenu() *GMenu {
	return gmenu{mnemonic: -1}
}
