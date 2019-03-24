// Package layouts handles all panel related widgets.
package layouts

/*
Dialog widget is the representation of a form's dialog.
	Parameters:
		text        : The title is displayed in the title bar.
		icon        : An image to be displayed in the titlebar of the dialog.
		modal       : A modal dialog grabs all the input to the components behind the dialog from the user. Default = false.
		resizable   : Set whether the dialog can be resized. Default = false.
		closable    : Set whether the dialog can be closed by a button. Default = false.
		maximizable : Set whether the dialog can be maximized by a button. Default = false.
		iconifiable	: Set whether the dialog can be iconified by a button. Default = false.

	Dialog is similar to panel, but it has border and title, and you can drag it.
*/
type Dialog struct {
	text        string
	icon        icon
	modal       bool
	resizable   bool
	closable    bool
	maximizable bool
	iconifiable bool
	panel       Panel
	component   Components
}
