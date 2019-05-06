// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

// GMenuItem is a single selectable choice in a menu grouping.
type GMenuItem interface{}

/*
gmenuitem parameters:

		GComboBox	: @see gcombobox.
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
	GComboBox
	accelerator GKeystroke
	mnemonic    int
}

// NewMenuItem returns a menu item structure with default values.
func NewMenuItem() GMenuItem {
	return &gmenuitem{mnemonic: -1}
}
