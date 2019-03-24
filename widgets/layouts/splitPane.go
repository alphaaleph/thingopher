// Package layouts handles all panel related widgets.
package layouts

/*
SplitPane widget is the representation of a form's split pane.
	Parameters:
		orientation	: The splitpane orientation is either horizontal or vertical. Default = horizontal.
		divider		: The location of the divider. A negative value implies the divider should be reset to a value that attempts to honor the preferred size of the two components. Default = -1.

	SplitPane is used to divide two components. The following one has two text-areas aligned top to bottom.

	<splitpane orientation="vertical" divider="24">
  		<textarea text="Top"/>
  		<textarea text="Bottom" />
	</splitpane>
*/
type SplitPane struct {
	orientation Choice
	divider     int
	component   Component
}
