// Package input handles all widgets that are used to collect information.
package input

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GSlider is the exported implementation of the slider widget
type GSlider interface{}

/*
gslider widget is the representation of a form's data collection slider.

	Parameters:
		component	: name, enabled, visible, tooltip, property, i18n, width, height, colspan, rowspan, weightx,
						weighty, halign, and valign parameters from gcomponent.
		orientation	: Possible values are: horizontal, and vertical. Default = horizontal.
		minimum		: The slider's minimum value. Default = 0.
		maximum		: The slider's maximum value. Default = 100.
		value		: The value is always between the slider's minimum and maximum values, inclusive. By default, the
						value equals the minimum.
		unit		: The distance of the value change when using arrow buttons. Default = 5.
		block		: The distance of the value change when using page buttons. Default = 25.

	A slider lets the user graphically select a value by sliding a knob within a bounded interval, e.g. between 25 and
	75.

		<gslider minimum="25" maximum="75" value="50" orientation="vertical" />

*/
type gslider struct {
	component   w.GComponent
	orientation p.GOrientation
	minimum     uint8
	maximum     uint8
	value       uint8
	unit        uint8
	block       uint8
}

// GSlider returns a slider object with default values.
func GSlider() *GSlider {
	return gslider{orientation: p.Horizontal, maximum: 100, unit: 5, block: 25}
}
