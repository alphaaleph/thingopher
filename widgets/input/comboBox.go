// Package input handles all widgets that are used to collect information.
package input

/*
ComboBox widget is the representation of a form's data collection combo box.
	Parameters:
		icon		: he icon image that the combobox displays.
		selected	: The index of the currently selected choice, value -1 indicates a custom edited value in an editable box. Default = -1.

	ComboBox is a combination of a text field and drop-down list. This example has a default value, two choices, and it is editable.

	<combobox text="ComboBox">
  		<choice text="Choice" icon="image.gif" />
  		<choice text="Disabled" enabled="false" />
	</combobox>
*/
type ComboBox struct {
	icon      Icon
	selected  int
	textField textField
	component Component
}
