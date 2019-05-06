// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package layouts handles all panel related widgets.
package layouts

import w "ebonynaranja.com/thingopher/widgets"

// GSeparator is used to separate widgets into different areas.
type GSeparator interface{}

/*
gseparator parameters:

		GComponent	: @see gcomponent.

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
	w.GComponent
}

// NewSeparator returns a separator structure with default values.
func NewSeparator() GSeparator {
	return &gseparator{}
}
