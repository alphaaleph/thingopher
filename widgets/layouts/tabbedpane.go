// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTabbedPane provides a panel that can have multiple tabs.
type GTabbedPane interface{}

/*
gtabbedpane parameters:

		GComponent	: @see gcomponent.
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
	w.GComponent
	placement p.Placement
	selected  int
}

// NewTabbedPane returns a tabbed pane structure with default values.
func NewTabbedPane() GTabbedPane {
	return &gtabbedpane{placement: p.Top}
}
