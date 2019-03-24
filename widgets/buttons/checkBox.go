// Package buttons manage all related button widgets
package buttons

/*
CheckBox widget is the representation of a form's check box.
	Parameters:
		text		: The text string that the checkbox displays.
		icon		: The icon image that the checkbox displays.
		alignment	: The alignment of the text and image similar to label. Possible values are: left, center, and right. Default = center.
		mnemonic	: Specifies the index of underlined char and a key combination (Alt + char) which change the checkboxs's state and invokes the action listener. Default = -1.
		selected	: The state of the checkbox. True if the checkbox is selected, false if it's not. Default = false.
		group		: Identifies the radio button group if not null. Only one radio button at a time can be selected. User can set on a radio button, the selected button of the group will be set off (the group members is searched only in the same parent).

	The first checkbox has a short text and an icon, its state is selected. The second one is a selected radio button, and the last is deselected. A set of radio buttons identified by the group string id.

	<checkbox text="CheckBox" icon="image.gif" selected="true"/>
	<checkbox text="RadioButton-on" group="group" selected="true"/>
	<checkbox text="RadioButton" group="group"/>
*/
type CheckBox struct {
	text      string
	icon      Icon
	alignment Choice
	mnemonic  int
	selected  bool
	group     string
	component Component
}
