// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package items handles all widgets that can hold data.
package items

import (
	i "ebonynaranja.com/thingopher/widgets/input"
)

// GNode holds the data of single entry in a tree branch.
type GNode interface{}

/*
gnode parameters:

		GComboBox	: @see gcombobox.
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
	i.GComboBox
	selected bool
	expanded bool
}

// NewNode returns a node structure with default values.
func NewNode() GNode {
	return &gnode{selected: false, expanded: true}
}
