// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTabbedPane is the exported implementation of the tabbed pane widget
type GTabbedPane interface{}

/*
TabbedPane widget is the representation of a form's tabbed pane.
	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		placement 	: The placement for the tabs relative to the content. Possible values are: top, left, bottom,
						right, and stack. The stack placement arranges tabs so that they resemble a sidebar.
						Default = top.
		selected	: The index of the currently selected tab, and the visible content (the first index is 0).
						Value -1 means there is no selected tab.

	gtabbedpane contains tabs and components. The first tab (has an index equal to 0) is associated with the first
	component. This example tabbed pane has 3 tab and 3 components (text areas), tabs are on the left, and the second
	component is visible.

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
type gtabbedpane struct {
	component w.GComponent
	placement p.GPlacement
	selected  int
}

// GTabbedPane returns a tabbed pane object with default values.
func GTabbedPane() *GTabbedPane {
	return gtabbedpane{placement: p.Top}
}
