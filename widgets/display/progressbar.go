// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package display manage all non-input widgets.
package display

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

/*
GProgressBar is used to measured a predetermine amount of time, or work steps.

	ProressBar displays an integer value within a bounded interval, the following between 25 and 75. Its value is
	currently in the middle.

		{"gprogressbar":{"minimum":25,"maximum":75,"value":50,"orientation":"vertical"}}

	Properties:
		Orientation	: Possible values are: horizontal, and vertical. Default = horizontal.
		Minimum		: The progressbar's minimum value. Default = 0.
		Maximum		: The progressbar's maximum value. Default = 100.
		Value		: The value is always between the progressbar's minimum and maximum values, inclusive. By default,
						the value equals the minimum. Default = 0.
		GComponent	: @see GComponent.
*/
type GProgressBar struct {
	Orientation p.Orientation `json:"orientation"`
	Minimum     uint8         `json:"minimum"`
	Maximum     uint8         `json:"maximum"`
	Value       uint8         `json:"value"`
	w.GComponent
}

// DefaultGProgressBar returns a progress bar structure with default values.
func DefaultGProgressBar() *GProgressBar {
	return &GProgressBar{Orientation: p.Horizontal, Maximum: 100}
}
