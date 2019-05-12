// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines some of the types used by the widgets parameters.
package properties

// Type defines the various widget's control and look.
type Type uint8

/*
	TNormal is the default value of a widget's category.
	TDefault is used as a widget's category control.
	TCancel is used as a widget's category control.
	TLink changes the appereance of a button so that it resembles HTML link.
*/
const (
	TNormal  Type = iota // 0x00
	TDefault             // 0x01
	TCancel              // 0x02
	TLink                // 0x03
)

//Text returns the string representation of the widget's type.
func (t Type) Text() string {
	switch t {
	case TNormal:
		return "Normal"
	case TDefault:
		return "Default"
	case TCancel:
		return "Cancel"
	case TLink:
		return "Link"
	default:
		return ""
	}
}

//List returns a slice of type strings.
func (t Type) List() []string {
	return []string{TNormal.Text(), TDefault.Text(), TCancel.Text(), TLink.Text()}
}
