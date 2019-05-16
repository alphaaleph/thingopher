// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package properties defines some of the types used by the widgets parameters.
package properties

// Selection defines the various widget selections.
type Selection uint8

/*
	Single selects a single entry in a list.
	Interval selects multiple entries in a list non-consecutively.
	Multiple selects multiple entries in a list consecutively.
*/
const (
	Single   Selection = iota // 0x00
	Interval                  // 0x01
	Multiple                  // 0x02
)

//String return a string representation of one of the enumerations
func (sel Selection) String() string {
	return [...]string{"Single", "Interval", "Multiple"}[sel]
}

//List returns a slice of selection strings.
func (sel Selection) List() []string {
	return []string{Single.String(), Interval.String(), Multiple.String()}
}
