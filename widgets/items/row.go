// Package items handles all widgets that can hold data.
package items

/*
Row widget is the representation of a form's table row.
	Parameters:
		selected	: True if currently selected the entire row. Default = false.

	Table contains rows (and columns for the header), and row contains cells. The 'row' key identifies the list of rows. An entire row (not cell) is selectable in a table.
*/
type Row struct {
	selected bool
	cells    []Cell
}
