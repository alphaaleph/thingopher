// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GPanel supports the grouping of several widgets together.
type GPanel interface{}

/*
gpanel parameters:

		GComponent	: @see gcomponent.
		icon       	: The icon image that the dialog title displays.
		columns    	: Specifies the number of available cells in a row. Default 0 value specifies 1 row and unlimited
						cell columns. Default = 0.
		top        	: The border gap from the top. It specifies the space that the panel must leave. Default = 0.
		left       	: The blank space from the left. Default = 0.
		bottom     	: The blank space from the bottom. Default = 0.
		right      	: The blank space from the right. Default = 0.
		gap        	: The horizontal and vertical gap between components. Default = 0.
		border     	: Border painted if true. Default = false.
		scrollable	: If present and set to true, add scrollbars to the panel if its contents is bigger than the
						panel's bounds. Default = false.
		text       	: The text string that the dialog title displays.

	gpanel is a container with a layout manager similar to GridBagLayout. The following example is similar to
	BorderLayout. This panel has 5 components, the first row contains only the 'North' component, the second row has 3
	components, and the last has only the 'South' one. The extra space is distributed to the 2nd row and the 2nd column.
	The gap between the components and on the border are 4pt.

		<gpanel columns="3" gap="4" top="4" left="4" bottom="4" right="4">
  			<gtextfield text="North" colspan="3" />
  			<glabel text="East" />
  			<gtextarea text="Center" weightx="1" weighty="1" />
  			<glabel text="West" />
  			<gtextfield text="South" colspan="3" />
		</gpanel>
*/
type gpanel struct {
	w.GComponent
	icon       p.Icon
	columns    uint16
	top        uint8
	left       uint8
	bottom     uint8
	right      uint8
	gap        uint8
	border     bool
	scrollable bool
	text       string
}

// NewPanel returns a panel structure with default values.
func NewPanel() GPanel {
	return &gpanel{border: false, scrollable: false}
}
