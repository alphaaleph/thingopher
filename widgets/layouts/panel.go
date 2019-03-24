// Package layouts handles all panel related widgets.
package layouts

/*
Panel widget is the representation of a form's panel.
	Parameters:
		columns    	: Specifies the number of available cells in a row. Default 0 value specifies 1 row and unlimited cell columns. Default = 0.
		top        	: The border gap from the top. It specifies the space that the panel must leave. Default = 0.
		left       	: The blank space from the left. Default = 0.
		bottom     	: The blank space from the bottom. Default = 0.
		right      	: The blank space from the right. Default = 0.
		gap        	: The horizontal and vertical gap between components. Default = 0.
		text       	: The text string that the dialog title displays.
		icon       	: The icon image that the dialog title displays.
		border     	: Border painted if true. Default = false.
		scrollable	: f present and set to true, add scrollbars to the panel if its contents is bigger than the panel's bounds. Default = false.

	Panel is a container with a layout manager similar to GridBagLayout. The following example is similar to BorderLayout. This panel has 5 components, the first row contains only the 'North' component, the second row has 3 components, and the last has only the 'South' one. The extra space is distributed to the 2nd row and the 2nd column. The gap between the components and on the border are 4pt.

	<panel columns="3" gap="4" top="4" left="4" bottom="4" right="4">
  		<textfield text="North" colspan="3" />
  		<label text="East" />
  		<textarea text="Center" weightx="1" weighty="1" />
  		<label text="West" />
  		<textfield text="South" colspan="3" />
	</panel>
*/
type Panel struct {
	columns    int
	top        int
	left       int
	bottom     int
	right      int
	gap        int
	text       string
	icon       Icon
	border     bool
	scrollable bool
	component  Component
}
