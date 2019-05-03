// Package items handles all widgets that can hold data.
package items

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GNode is the exported implementation of the node widget
type GNode interface{}

/*
gnode widget is the representation of a form's data node.

	Parameters:
		combobox	: name, enabled, text, icon, alignment, tooltip, and property parameters similar to gcombobox.
		selected	: True if currently selected. Default = false.
		expanded 	: Ensures that the node is expanded if true, otherwise collapsed. Default = true.

	Tree node is similar to list item, but maybe contains subnodes, and has collapse control.

		<gtree selection="multiple">
  			<gnode text="Node A" icon="image.gif">
    			<gnode text="Node B" icon="image.gif" selected="true" />
    			<gnode text="Node C" icon="image.gif" />
  			</gnode>
  			<gnode text="Node D" expanded="false">
    			gnode text="Node E" icon="image.gif" />
  			</gnode>
		</gtree>
*/
type gnode struct {
	combobox i.GComboBox
	selected bool
	expanded bool
}

// GNode returns a node object with default values.
func GNode() *GNode {
	return gtree{selected: false, expanded: true}
}
