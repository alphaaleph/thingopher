// Package buttons manage all related button widgets
package buttons

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GButton interface exported instead of the struct
type GButton interface{}

/*
gbutton widget is the representation of a form's button.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from component.
		icon		: The icon image that the checkbox displays.
		alignment	: The alignment of the text and image similar to label. Possible values are: center, left, and
						right. The default value is center. Icons are always displayed to the left of text (if any).
						Default = center.
		category	: Possible values are: normal, default, cancel, and link. The default value is normal. Default, and
						cancel values are for dialog control. Link changes the appearance of button so that it
						resembles HTML link. Default = normal.
		mnemonic 	: Specifies the index of underlined char and a key combination (Alt + the char) which invokes the
						(not necessarily focused, but enabled and visible) button's action listener. Default = -1.
		text		: The text string that the button displays.

	The following button displays a short text string and an image. Its content is on the left and centered vertically.
	The second button resembles an HTML link, and when pressed invokes the method goHome.

		<gbutton text="Button" icon="image.gif" alignment="left" tooltip="ToolTip" />
		<gbutton text="www.thinlet.com" type="link" action="goHome" />
*/
type gbutton struct {
	component w.Component
	icon      p.Icon
	alignment p.Alignment
	category  p.Category
	mnemonic  int
	text      string
}

// GButton returns a button object with default values.
func GButton() *GButton {
	return gbutton{alignment: p.ACenter, category: p.CNormal, mnemonic: -1}
}
