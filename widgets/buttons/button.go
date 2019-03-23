// Package buttons manage all related button widgets
package buttons

/*
Button widget is the representation of a form's button.
	Parameters:
		text		: The text string that the button displays.
		icon		: The icon image that the button displays.
		alignment	: The alignment of the text and image similar to label. Possible values are: center, left, and right. The default value is center. Icons are always displayed to the left of text (if any). Default = center.
		mnemonic	: Specifies the index of underlined char and a key combination (Alt + the char) which invokes the (not necessarily focused, but enabled and visible) button's action listener. Default = -1.
		type_		: Possible values are: normal, default, cancel, and link. The default value is normal. Default, and cancel values are for dialog control. Link changes the appearance of button so that it resembles HTML link. Default = normal.

	The following button displays a short text string and an image. Its content is on the left and centered vertically. The second button resembles an HTML link, and when pressed invokes the method goHome.

	<button text="Button" icon="image.gif" alignment="left" tooltip="ToolTip" />
	<button text="www.thinlet.com" type="link" action="goHome" />
*/
type Button struct {
	text      string
	icon      Icon
	alignment Choice
	mnemonic  int
	type_     Choice
	*Component
}

// LISTENERS

// Action invokes the given method when the button is pressed by mouse or keyboard.
func (b *Button) Action() {
	panic("Not implemented yet")
}

// KEYBOARD

// SpaceBar is used to activate the button.
func (b *Button) SpaceBar() {
	panic("Not implemented yet")
}

// Tab is used to navigate forward.
func (b *Button) Tab() {
	panic("Not implemented yet")
}

// ShiftTab is used to navigate backwards.
func (b *Button) ShiftTab() {
	panic("Not implemented yet")
}
