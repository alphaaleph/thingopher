// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

// GMenu provides the management and selection of a single set of menu items, in a menu bar.
type GMenu interface{}

/*
gmenu parameters:

		GComboBox	: @see gcombobox.
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
	GComboBox
	mnemonic int
}

// NewMenu returns a menu structure with default values.
func NewMenu() GMenu {
	return &gmenu{mnemonic: -1}
}
