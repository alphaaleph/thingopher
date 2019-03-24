// Package display manage all non-input widgets.
package display

/*
Label widget is the representation of a form's label.
	Parameters:
		text		: The text string that the label displays.
		icon		: The icon image that the label displays.
		alignment	: The alignment of the label's content along the X axis, vertically is centered. Possible values are: left, center, and right. The default value, if not set, is left. The image position horizontally left, above, and right relative to the text. Default = left.
		mnemonic	: Specifies the underlined char in the label's text. Default = -1.
		for_		: Set the component this is labelling. The label will focus the component specified by its name when the mnemonic is activated.

	The following label displays a short text string and an image. Its content is centered both vertically and horizontally.

	<label text="Label" icon="image.gif" alignment="center" />
*/
type Label struct {
	text      string
	icon      Icon
	alignment Choice
	mnemonic  int
	for_      *Component
	component Component
}
