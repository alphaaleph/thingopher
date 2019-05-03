// Package items handles all widgets that can hold data.
package items

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTree is the exported implementation of the tree widget
type GTree interface{}

/*
gtree widget is the representation of a form's data tree.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		list		: line, and selection parameters from list.
		selection	: Possible values are: single, interval, and multiple. For the default single value the selection
						can only contain one path at a time, for interval the selection can only be contiguous (of the
						currently visible nodes), and for the multiple value the selection can contain any number of
						nodes that are not necessarily contiguous. Default = single.
		angle     	: To specify to draw lines detailing the relationships between nodes. Default false.

	Tree displays a set of hierarchical data, contains nodes (not only one root node is allowed), and a node could have
	subnodes. It internally handles scrolling. The following example shows a tree with 2 root nodes ('Node A', and
	'Node D'), the first hase.

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
type gtree struct {
	component w.GComponent
	list      GList
	selection p.GSelection
	angle     bool
}

// GTree returns a tabbed pane object with default values.
func GTree() *GTree {
	return gtree{selection: p.Single, angle: false}
}
