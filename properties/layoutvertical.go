// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines types used by the widgets parameters.
package properties

// LayoutVertical defines the various widget vertical aligments.
type LayoutVertical uint8

/*
	VFill indicates that the component will use all the space.
	VCenter indicates that the component will centered in the total space.
	VTop indicates that the component will be aligned on the top side.
	VBottom indicates that the component will be aligned on the bottom side.
*/
const (
	VFill   LayoutVertical = iota // 0x00
	VCenter                       // 0x01
	VTop                          // 0x02
	VBottom                       // 0x03
)

//Text returns the string representation of the vertical layout.
func (lv LayoutVertical) Text() string {
	switch lv {
	case VFill:
		return "Fill"
	case VCenter:
		return "Center"
	case VTop:
		return "Top"
	case VBottom:
		return "Bottom"
	default:
		return ""
	}
}

//List returns a slice of vertical layout strings.
func (lv LayoutVertical) List() []string {
	return []string{VFill.Text(), VCenter.Text(), VTop.Text(), VBottom.Text()}
}
