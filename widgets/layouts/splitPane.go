// Package layouts handles all panel related widgets.
package layouts

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GSplitPane is the exported implementation of the separator widget
type GSplitPane interface{}

/*
gsplitpane widget is the representation of a form's split pane.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
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
	component   w.GComponent
	orientation p.GOrientation
	divider     int
}

// GSplitPane returns a panel object with default values.
func GSplitPane() *GSplitPane {
	return gsplitpane{orientation: p.Horizontal, divider: -1}
}
