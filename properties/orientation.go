// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines some of the types used by the widgets parameters.
package properties

// Orientation defines the various widgets orientations.
type Orientation uint8

/*
	Horizontal indicates that a widget will be displayed in horizontal fashion.
	Vertical indicates that a widget will be displayed in vertical fashion.
*/
const (
	Horizontal Orientation = iota // 0x00
	Vertical                      // 0x01
)

//Text returns the string representation of the orientation choice.
func (orn Orientation) Text() string {
	switch orn {
	case Horizontal:
		return "Horizontal"
	case Vertical:
		return "Vertical"
	default:
		return ""
	}
}

//List returns a slice of orientation strings.
func (orn Orientation) List() []string {
	return []string{Horizontal.Text(), Vertical.Text()}
}
