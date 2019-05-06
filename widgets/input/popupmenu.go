// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package input handles all widgets that are used to collect information.
package input

import (
	w "ebonynaranja.com/thingopher/widgets"
)

// GPopupMenu resemble in functionality normal menus, and can be added to most components. GPopupMenu component can
// have the same sub-components as menu. Popup menus are invoked with mouse right-click.
type GPopupMenu interface{}

/*
gpopupmenu parameters:

	GBaseComponent	: @see gbasecomponent.
*/
type gpopupmenu struct {
	w.GBaseComponent
}

// NewPopupMenu returns a popup menu structure with default values.
func NewPopupMenu() GPopupMenu {
	return &gpopupmenu{}
}
