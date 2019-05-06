// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines some of the types used by the widgets parameters.
package properties

// Sort defines the various widget sorts.
type Sort uint8

/*
	None indicates that the column is not sorted.
	Ascent indicates that the column will be sorted in ascending order.
	Descent indicates that the column will be sorted in descending order.
*/
const (
	None    Sort = iota // 0x00
	Ascent              // 0x01
	Descent             // 0x02
)

//Text returns the string representation of the selection choice.
func (st Sort) Text() string {
	switch st {
	case None:
		return "None"
	case Ascent:
		return "Ascent"
	case Descent:
		return "Descent"
	default:
		return ""
	}
}

//List returns a slice of sort strings.
func (st Sort) List() []string {
	return []string{None.Text(), Ascent.Text(), Descent.Text()}
}
