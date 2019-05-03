// Package input handles all widgets that are used to collect information.
package input

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GCheckBoxMenuItem is the exported implementation of the check box menu item widget
type GCheckBoxMenuItem interface{}

/*
gcheckboxmenuitem widget is the representation of a form's data collection check box menu item.

	Parameters:
		menuitem	: mnemonic, accelerator, and action parameters described at menuitem.
		combobox	: name, enabled, text, icon, alignment, tooltip, and property parameters similar to gcombobox.
		selected	: The state of the checkboxmenuitem. Default = false.
		group		: Identifies a group if not null. Only one radio menuitem at a time can be selected in a group.
						User can set on a radio menuitem, the selected menuitem of the group will be set off (the group
						members is searched only in the same menu).

	A menu item that can be selected or deselected. The selected menuitem appears with a checkmark next to it. A
	radio-button menu-item is a menu item that is part of a group of menu items in which only one item in the group can
	be selected.

		<gpanel columns="1" gap="4">
  			<gcheckboxmenuitem weightx="1">
    			<gmenu text="Menu">
      				<gmenuitem text="MenuItem" icon="image.gif" />
      				<gmenu text="Menu">
        				<gcheckboxmenuitem text="CheckBox" selected="true" />
        				<gseparator />
        				<gcheckboxmenuitem text="RadioButton" group="group"/>
      				</gmenu>
    			</gmenu>
  			</gcheckboxmenuitem>
  			<gtextarea weighty="1" />
		</gpanel>
*/
type gcheckboxmenuitem struct {
	menuitem GMenuItem
	combobox i.GComboBox
	selected bool
	group    string
}

// GCheckBoxMenuItem returns a check box menu item object with default values.
func GCheckBoxMenuItem() *GCheckBoxMenuItem {
	return gcheckboxmenuitem{selected: false}
}
