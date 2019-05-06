// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GSplitPane allows to divide components.
type GSplitPane interface{}

/*
gsplitpane parameters:

		GComponent	: @see gcomponent.
		orientation	: The splitpane orientation is either horizontal or vertical. Default = horizontal.
		divider		: The location of the divider. A negative value implies the divider should be reset to a value that
						attempts to honor the preferred size of the two components. Default = -1.

	gsplitpane is used to divide two components. The following one has two text-areas aligned top to bottom.

		<gsplitpane orientation="vertical" divider="24">
  			<gtextarea text="Top"/>
  			<gtextarea text="Bottom" />
		</gsplitpane>
*/
type gsplitpane struct {
	w.GComponent
	orientation p.Orientation
	divider     int
}

// NewSplitPane returns a split panel structure with default values.
func NewSplitPane() GSplitPane {
	return &gsplitpane{orientation: p.Horizontal, divider: -1}
}
