// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GTree is a hierarchical outline view of data.
type GTree interface{}

/*
gtree parameters:

		GComponent	: @see gcomponent.
		GList		: @see glist.
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
	w.GComponent
	GList
	selection p.Selection
	angle     bool
}

// NewTree returns a tree structure with default values.
func NewTree() GTree {
	return &gtree{selection: p.Single, angle: false}
}
