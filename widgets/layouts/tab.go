// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GTab manages a single tab from a tabbed pane control.
type GTab interface{}

/*
gtab parameters:

		GComboBox	: @see gcombobox.
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
	i.GComboBox
	mnemonic int
}

// NewTab returns a tab structure with default values.
func NewTab() GTab {
	return &gtab{mnemonic: -1}
}
