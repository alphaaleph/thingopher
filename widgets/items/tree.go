// Package items handles all widgets that can hold data.
package items

/*
Tree widget is the representation of a form's data tree.
	Parameters:
		selection	: Possible values are: single, interval, and multiple. For the default single value the selection can only contain one path at a time, for interval the selection can only be contiguous (of the currently visible nodes), and for the multiple value the selection can contain any number of nodes that are not necessarily contiguous. Default = single.
		angle     	: To specify to draw lines detailing the relationships between nodes. Default false.

	Tree displays a set of hierarchical data, contains nodes (not only one root node is allowed), and a node could have subnodes. It internally handles scrolling. The following example shows a tree with 2 root nodes ('Node A', and 'Node D'), the first hase.

	<tree selection="multiple">
  		<node text="Node A" icon="image.gif">
    		<node text="Node B" icon="image.gif" selected="true" />
    		<node text="Node C" icon="image.gif" />
  		</node>
  		<node text="Node D" expanded="false">
    		<node text="Node E" icon="image.gif" />
  		</node>
	</tree>
*/
type Tree struct {
	selection Choice
	angle     bool
	list      List
	component Component
}
