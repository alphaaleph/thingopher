// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines types used by the widgets parameters.
package properties

// LayoutHorizontal defines the various widget horizontal aligments.
type LayoutHorizontal uint8

/*
	HFill indicates that the component will use all the space.
	HCenter indicates that the component will centered in the total space.
	HLeft indicates that the component will be aligned on the left side.
	HRight indicates that the component will be aligned on the right side.
*/
const (
	HFill   LayoutHorizontal = iota // 0x00
	HCenter                         // 0x01
	HLeft                           // 0x02
	HRight                          // 0x03
)

//Text returns the string representation of the horizontal layout.
func (lh LayoutHorizontal) Text() string {
	switch lh {
	case HFill:
		return "Fill"
	case HCenter:
		return "Center"
	case HLeft:
		return "Left"
	case HRight:
		return "Right"
	default:
		return ""
	}
}

//List returns a slice of horizontal layout strings.
func (lh LayoutHorizontal) List() []string {
	return []string{HFill.Text(), HCenter.Text(), HLeft.Text(), HRight.Text()}
}
