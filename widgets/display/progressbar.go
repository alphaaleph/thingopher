// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package display manage all non-input widgets.
package display

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GProgressBar is used to measured a predetermine amount of time, or work steps.
type GProgressBar interface{}

/*
gprogressbar parameters:

		GComponent	: @see gcomponent.
		orientation	: Possible values are: horizontal, and vertical. Default = horizontal.
		minimum		: The progressbar's minimum value. Default = 0.
		maximum		: The progressbar's maximum value. Default = 100.
		value		: The value is always between the progressbar's minimum and maximum values, inclusive. By default,
						the value equals the minimum. Default = 0.

	ProressBar displays an integer value within a bounded interval, the following between 25 and 75. Its value is
	currently in the middle.

		<gprogressbar minimum="25" maximum="75" value="50" orientation="vertical" />
*/
type gprogressbar struct {
	w.GComponent
	orientation p.Orientation
	minimum     uint8
	maximum     uint8
	value       uint8
}

// NewProgressBar returns a progress bar structure with default values.
func NewProgressBar() GProgressBar {
	return &gprogressbar{orientation: p.Horizontal, maximum: 100}
}
