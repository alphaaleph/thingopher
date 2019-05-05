// Package widgets holds all the components supported by Thinggopher
package widgets

import (
	"image/color"

	"golang.org/x/image/font"

	p "ebonynaranja.com/thingopher/properties"
)

/*
gcomponent is a parent structure for most Thingopher widgets. All components support the parameters listed below.

	Parameters:
		property	: Binds an arbitrary key/value client property (or properties).
		halign		: Horizontal alignment of the component when more space available in the cell. Possible values
						are: fill, center, left, and right. Default = fill.
		valign		: Vertical alignment in the cell. Possible values are: fill, center, top, and bottom.
						Default = fill.
		background	: Custom background color to use, instead of the default background color. Note: for components
						with gradients (buttons, menubars, etc...) this means that the background will be filled with
						solid color, and not a custom gradient.
		foreground	: Custom foreground (text) color to use, instead of the default text color.
		font		: Custom font for texts.
		height		: Preferred height of the component, or 0. Default = 0.
		width		: Fixed preferred width of the component irrespectively of its content. If it is 0, the component
						will be asked for the preferred width. Default = 0.
		colspan		: Specifies the number of cells in a column in the component's display area. Default = 1.
		rowspan		: Specifies the number of cells in a row occupied by the component. Default = 1.
		weightx		: Used to determine how to distribute horizontal space among grid cells when more space is
						available for its parent component than required. Default = 0.
		weighty		: The extra vertical space is distributed to each cell height in proportion to its weight. Default
						value is 0. The row contains 0 weighted components remains the calculated preferred size.
						Default = 0.
		enabled 	: Enables or disables this component. A disabled component is painted gray, and can't respond to
						user mouse or keyboard input, gain the keyboard focus, and generate events. Default = true.
		visble		: An invisible component doesn't take place in parent's layout. Default = true.
		name		: Identifies the component.
		tooltip		: The text pops up when the mouse lingers inside the component.
*/
type gcomponent struct {
	property   p.Property
	halign     p.ChoiceHorizontal
	valign     p.ChoiceVertical
	background color.Color
	foreground color.Color
	font       font.Style
	height     uint16
	width      uint16
	colspan    uint16
	rowspan    uint16
	weightx    uint16
	weighty    uint16
	enabled    bool
	visible    bool
	name       string
	tooltip    string
}

// GComponent returns a component object with default values.
func GComponent() *GComponent {
	return gcomponent{colspan: 1, rowspan: 1, enabled: true, visible: true}
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
