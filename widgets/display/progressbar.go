// Package display manage all non-input widgets.
package display

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GProgressBar is the exported implementation of the progress bar widget
type GProgressBar interface{}

/*
gprogressbar widget is the representation of a form's progress bar.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from component.
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
	component   w.GComponent
	orientation p.GOrientation
	minimum     uint8
	maximum     uint8
	value       uint8
}

// GProgressBar returns a progress bar object with default values.
func GProgressBar() *GProgressBar {
	return gprogressbar{orientation: p.Horizontal, maximum: 100}
}
