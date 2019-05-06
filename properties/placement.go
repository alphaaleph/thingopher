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

//Text returns the string representation of the placement choice.
func (pl Placement) Text() string {
	switch pl {
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	default:
		return ""
	}
}

//List returns a slice of placement strings.
func (pl Placement) List() []string {
	return []string{Top.Text(), Bottom.Text()}
}
