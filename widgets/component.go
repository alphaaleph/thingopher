// Package widgets holds all the components supported by Thinggopher
package widgets

import (
	"image/color"

	"golang.org/x/image/font"

	"theasilum.com/thingopher/properties"
)

/*
Component is a parent structure for most Thingopher widgets. All components support the parameters listed below.
	Parameters:
		name		: Identifies the component.
		enabled 	: Enables or disables this component. A disabled component is painted gray, and can't respond to user mouse or keyboard input, gain the keyboard focus, and generate events. Default = true.
		visble		: An invisible component doesn't take place in parent's layout. Default = true.
		tooltip		: The text pops up when the mouse lingers inside the component.
		font		: Custom font for texts.
		foreground	: Custom foreground (text) color to use, instead of the default text color.
		background	: Custom background color to use, instead of the default background color. Note: for components with gradients (buttons, menubars, etc...) this means that the background will be filled with solid color, and not a custom gradient.
		width		: Fixed preferred width of the component irrespectively of its content. If it is 0, the component will be asked for the preferred width. Default = 0.
		height		: Preferred height of the component, or 0. Default = 0.
		colspan		: Specifies the number of cells in a column in the component's display area. Default = 1.
		rowspan		: Specifies the number of cells in a row occupied by the component. Default = 1.
		weightx		: Used to determine how to distribute horizontal space among grid cells when more space is available for its parent component than required. Default = 0.
		weighty		: The extra vertical space is distributed to each cell height in proportion to its weight. Default value is 0. The row contains 0 weighted components remains the calculated preferred size. Default = 0.
		halign		: Horizontal alignment of the component when more space available in the cell. Possible values are: fill, center, left, and right. Default = fill.
		valign		: Vertical alignment in the cell. Possible values are: fill, center, top, and bottom. Default = fill.
		property	: Binds an arbitrary key/value client property (or properties).
*/
type Component struct {
	name       string
	enabled    bool
	visible    bool
	tooltip    string
	font       font.Style
	foreground color.Color
	background color.Color
	width      int
	height     int
	colspan    int
	rowspan    int
	weightx    int
	weighty    int
	halign     properties.ChoiceHorizontal
	valign     properties.ChoiceVertical
	property   properties.Property
}

// LISTENERS

// Init is a method to invoke only once when the loading of the xml resource (including this component) is finished.
func (component *Component) Init() {
	panic("Not implemented yet.")
}

// FocusLost invoked when a component loses the keyboard focus, thus it is no longer the focus owner.
func (component *Component) FocusLost() {
	panic("Not implemented yet.")
}

// FocusGained invoked when a component gains the keyboard focus, thus it is now the focus owner.
func (component *Component) FocusGained() {
	panic("Not implemented yet.")
}
