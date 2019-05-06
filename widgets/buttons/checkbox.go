// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

// GCheckBox supports the use of single/multiple choice.
type GCheckBox interface{}

/*
gcheckbox parameters:

		GComponent	: @see gcomponent.
		icon		: The icon image that the checkbox displays.
		mnemonic	: Specifies the index of underlined char and a key combination (Alt + char) which change the
						checkboxs's state and invokes the action listener. Default = -1.
		alignment	: The alignment of the text and image similar to label. Possible values are: left, center, and
						right. Default = center.
		selected	: The state of the checkbox. True if the checkbox is selected, false if it's not. Default = false.
		group		: Identifies the radio button group if not null. Only one radio button at a time can be selected.
						User can set on a radio button, the selected button of the group will be set off (the group
						members is searched only in the same parent).
		text		: The text string that the checkbox displays.

	The first checkbox has a short text and an icon, its state is selected. The second one is a selected radio button,
	and the last is deselected. A set of radio buttons identified by the group string id.

		<gcheckbox text="CheckBox" icon="image.gif" selected="true"/>
		<gcheckbox text="RadioButton-on" group="group" selected="true"/>
		<gcheckbox text="RadioButton" group="group"/>
*/
type gcheckbox struct {
	component w.GComponent
	icon      p.Icon
	mnemonic  int
	alignment p.Alignment
	selected  bool
	group     string
	text      string
}

// NewCheckBox returns a check box structure with default values.
func NewCheckBox() GCheckBox {
	return &gcheckbox{selected: false, alignment: p.ACenter, mnemonic: -1}
}
