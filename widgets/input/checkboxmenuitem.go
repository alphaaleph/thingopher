// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

// GCheckBoxMenuItem is a menu item that allows to be checked/unchecked.
type GCheckBoxMenuItem interface{}

/*
gcheckboxmenuitem parameters:

		GMenuItem	: @see gmenuitem.
		GComboBox	: @see gcombobox.
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
	GMenuItem
	GComboBox
	selected bool
	group    string
}

// NewCheckBoxMenuItem returns a check box menu item structure with default values.
func NewCheckBoxMenuItem() GCheckBoxMenuItem {
	return &gcheckboxmenuitem{selected: false}
}
