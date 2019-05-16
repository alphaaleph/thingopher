// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines some of the types used by the widgets parameters.
package properties

// Placement defines the various widgets placements.
type Placement uint8

/*
	Top indicates that the menu bar will be placed at the top of a container.
	Bottom indicates that the menu bar will be placed at the bottom of a container.
*/
const (
	Top    Placement = iota // 0x00
	Bottom                  // 0x01
)

//String return a string representation of one of the enumerations
func (pl Placement) String() string {
	return [...]string{"Top", "Bottom"}[pl]
}

//List returns a slice of placement strings.
func (pl Placement) List() []string {
	return []string{Top.String(), Bottom.String()}
}
