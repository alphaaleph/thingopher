// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package widgets holds all the components supported by thingopher
package widgets

import (
	"image/color"

	p "ebonynaranja.com/thingopher/properties"
	"golang.org/x/image/font"
)

// GComponent is the parent structure for most thingopher widgets.
type GComponent interface{}

/*
gcomponent parameters:

		GBaseComponent	: @see gbasecomponent.
		background		: Custom background color to use, instead of the default background color. Note: for components
							with gradients (buttons, menubars, etc...) this means that the background will be filled
							with solid color, and not a custom gradient.
		foreground		: Custom foreground (text) color to use, instead of the default text color.
		font			: Custom font for texts.
		height			: Preferred height of the component, or 0. Default = 0.
		width			: Fixed preferred width of the component irrespectively of its content. If it is 0, the
							component will be asked for the preferred width. Default = 0.
		colspan			: Specifies the number of cells in a column in the component's display area. Default = 1.
		rowspan			: Specifies the number of cells in a row occupied by the component. Default = 1.
		weightx			: Used to determine how to distribute horizontal space among grid cells when more space is
							available for its parent component than required. Default = 0.
		weighty			: The extra vertical space is distributed to each cell height in proportion to its weight.
							Default value is 0. The row contains 0 weighted components remains the calculated preferred
							size. Default = 0.
		halign			: Horizontal alignment of the component when more space available in the cell. Possible values
							are: fill, center, left, and right. Default = fill.
		valign			: Vertical alignment in the cell. Possible values are: fill, center, top, and bottom.
							Default = fill.
*/
type gcomponent struct {
	GBaseComponent
	background color.Color
	foreground color.Color
	font       font.Style
	height     uint16
	width      uint16
	colspan    uint16
	rowspan    uint16
	weightx    uint16
	weighty    uint16
	halign     p.LayoutHorizontal
	valign     p.LayoutVertical
}

// NewComponent returns a component structure with default values.
func NewComponent() GComponent {
	return &gcomponent{colspan: 1, rowspan: 1}
}
