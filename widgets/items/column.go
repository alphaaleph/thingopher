// Package items handles all widgets that can hold data.
package items

/*
Column widget is the representation of a form's table column.
	Parameters:
		width	: The width of the column. The last (rightmost) column always extends to fill in the remaining width of the table, as determined by the parent container. Default = 80.
		sort	: If present, and set to either "ascent" or "descent", an arrow is drawn near right end of the column, as a visual indication of sorting order. NOTE: for now this is just a visual decoration, application still needs to do the sorting. Default = none.

	Table header includes the columns. They define the column texts and widths. If the total sum of column widths exceeds table bounds set by a parent container, horizontal scrollbar is drawn.
*/
type Column struct {
	width int
	sort  Choice
	label Label
}
