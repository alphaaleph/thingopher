// Package items handles all widgets that can hold data.
package items

/*
Item widget is the representation of a form's data item.
	Parameters:
		selected : True if the item currently selected. Default = true.

	List contains items, items display a short text string and an image and can be selected.
*/
type Item struct {
	selected bool
	combobox ComboBox
}
