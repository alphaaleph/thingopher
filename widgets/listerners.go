// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package widgets holds all the components supported by thingopher
package widgets

// IListener provides methods for the widgets to execute.
type IListener interface {
	Action()
}
