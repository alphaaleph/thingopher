// Package buttons manage all related button widgets
package buttons

/*
ToggleButton widget is the representation of a form's toggle button.
	Parameters:
		text		: The text string that the checkbox displays.
		icon		: The icon image that the checkbox displays.
		alignment	: The alignment of the text and image similar to label. Possible values are: left, center, and right. Default = center.
		mnemonic	: Specifies the index of underlined char and a key combination (Alt + char) which change the checkboxs's state and invokes the action listener. Default = -1.
		selected	: The state of the checkbox. True if the checkbox is selected, false if it's not. Default = false.

	An implementation of a two-state button, behaves as checkbox.

	<togglebutton text="ToggleButton" icon="image.gif" selected="true"/>
	<togglebutton text="ToggleButton" group="group"/>
*/
type ToggleButton struct {
	text      string
	icon      Icon
	alignment Choice
	mnemonic  int
	selected  bool
	component Component
}
