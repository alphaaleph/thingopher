// Package layouts handles all panel related widgets.
package layouts

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GTab is the exported implementation of the tab widget
type GTab interface{}

/*
gtab widget is the representation of a form's tab.

	Parameters:
		combobox	: name, enabled, text, icon, alignment, tooltip, and property parameters similar to gcombobox.
		mnemonic	: Specifies the underlined char and the key combination which selects the tab. Default = -1.

	Tabs are identified in the component's list as 'gtab', and components with the 'gcomponent' key.

		<gtabbedpane placement="left" selected="1" action="tabchanged">
  			<gtab text="One" icon="image.gif">
    			<gtextarea text="One" />
  			</gtab>
  			<gtab text="Two" alignment="right">
    			<gtextarea text="Two" />
  			</gtab>
  			<gtab text="Three" enabled="false">
    			<gtextarea text="Three" />
  			</gtab>
		</gtabbedpane>
*/
type gtab struct {
	combobox i.GComboBox
	mnemonic int
}

// GTab returns a tab object with default values.
func GTab() *GTab {
	return gtab{mnemonic: -1}
}
