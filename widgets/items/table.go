// Package items handles all widgets that can hold data.
package items

/*
Table widget is the representation of a form's data table.
	Parameters:
		selection	: Possible values are: single, interval, and multiple. The default single value allows to select one row at a time, the interval value to select one contiguous range of rows, and the multiple value to select one or more contiguous ranges of rows. Default = single.

	A table presents data in a two-dimensional table format, allows to select rows, has header, and internally handles scrolling. The following example table includes two columns ('A', and 'B'), and two rows, the first row is selected, but allows to select multiple rows.

	<table selection="multiple">
  		<header>
    		<column text="A" width="24" /><column text="B" />
  		</header>
  		<row selected="true">
    		<cell text="a1" /><cell text="b1" enabled="false" />
  		</row>
  		<row>
    		cell text="a2" /><cell text="b2" />
  		</row>
	</table>
*/
type Table struct {
	selection Choice
	line      Line
	component Component
}
