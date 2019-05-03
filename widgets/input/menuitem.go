// Package input handles all widgets that are used to collect information.
package input

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GMenuItem is the exported implementation of the menu item widget
type GMenuItem interface{}

/*
gmenuitem widget is the representation of a form's data collection menu item.

	Parameters:
		combobox	: name, enabled, text, icon, alignment, tooltip, and property parameters similar to gcombobox.
		accelerator	: The key combination which invokes the menuitem's action method without navigating the menu
						hierarchy.
		mnemonic	: Specifies the underlined char and the key combination which selects the tab. Default = -1.

	A regular menu item can have either text or a graphic icon associated with it, or both.

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
type gmenuitem struct {
	combobox    i.GComboBox
	accelerator GKeystroke
	mnemonic    int
}

// GMenuItem returns a menu item object with default values.
func GMenuItem() *GMenuItem {
	return gmenuitem{mnemonic: -1}
}
