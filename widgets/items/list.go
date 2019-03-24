// Package items handles all widgets that can hold data.
package items

/*
List widget is the representation of a form's data list.
	Parameters:
		selection	: Possible values are: single, interval, and multiple. The default single value allows to select one list item at a time, the interval value to select one contiguous range of items, and the multiple value to select one or more contiguous ranges of items. Default = single.
		line	  	: If present and set to false, don't draw lines separating the list items. Default = true.

	A list allows the user to select one or more objects from a list, and internally handles scrolling. The following example list has three items, the first is selected and the last is disabled, allows to select multiple items.

	<list selection="multiple">
  		<item text="List" selected="true" />
  		<item text="Item"/>
  		<item text="Disabled" enabled="false" />
	</list>
*/
type List struct {
	selection Choice
	line      bool
	component Component
}
