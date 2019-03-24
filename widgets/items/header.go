// Package items handles all widgets that can hold data.
package items

/*
Header widget is the representation of a form's table header.
	Parameters:
		columns	:

	Table header includes the columns. They define the column texts and widths.
*/
type Header struct {
	columns []Column
}
