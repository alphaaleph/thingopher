// Package layouts handles all panel related widgets.
package layouts

import w "ebonynaranja.com/thingopher/widgets"

// GSeparator is the exported implementation of the separator widget
type GSeparator interface{}

/*
gseparator widget is the representation of a form's separators.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.

	Separator is a horizontal or a vertical line. The following panel has a button('Left'), a vertical separator line,
	and a second button ('Right') in the first row, a horizontal separator in the second row, and the final button
	('Bottom') in the third row.

		<gpanel columns="3" gap="4" top="4" left="4">
  			<gbutton text="Left" />
  			<gseparator />
  			<gbutton text="Right" />
  			<gseparator colspan="3" />
  			<gbutton text="Bottom" colspan="3" />
		</gpanel>
*/
type gseparator struct {
	component w.GComponent
}

// GSeparator returns a panel object with default values.
func GSeparator() *GSeparator {
	return gseparator{}
}
