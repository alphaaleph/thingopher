// Package items handles all widgets that can hold data.
package items

/*
Node widget is the representation of a form's data node.
	Parameters:
		selected	: True if currently selected. Default = false.
		expanded 	: Ensures that the node is expanded if true, otherwise collapsed. Default = true.

	Tree node is similar to list item, but maybe contains subnodes, and has collapse control.
*/
type Node struct {
	selected bool
	expanded bool
	combobox ComboBox
}
