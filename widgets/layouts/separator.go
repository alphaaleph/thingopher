// Package layouts handles all panel related widgets.
package layouts

/*
Separator widget is the representation of a form's separators.
	Parameters:

	Separator is a horizontal or a vertical line. The following panel has a button('Left'), a vertical separator line, and a second button ('Right') in the first row, a horizontal separator in the second row, and the final button ('Bottom') in the third row.

	<panel columns="3" gap="4" top="4" left="4">
  		<button text="Left" />
  		<separator />
  		<button text="Right" />
  		<separator colspan="3" />
  		<button text="Bottom" colspan="3" />
	</panel>
*/
type Separator struct {
	component Component
}
