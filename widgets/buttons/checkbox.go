// Copyright 2019 EbonyNaranja. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

// Package buttons manage all related button widgets
package buttons

import (
	"fmt"

	p "ebonynaranja.com/thingopher/properties"
	w "ebonynaranja.com/thingopher/widgets"
)

/*
GCheckBox supports the use of single/multiple choice.

	The first checkbox has a short text and an icon, its state is selected. The second one is a selected radio button,
	and the last is deselected. A set of radio buttons identified by the group string id.

		{"gcheckbox":{"text":"CheckBox","icon":"image.gif","selected"="true"}}
		{"gcheckbox":{"text":"RadioButton-on","group":"group","selected"="true"}}
		{"gcheckbox":{"text":"RadioButton","group":"group"}}

	Properties:
		Icon		: The icon image that the checkbox displays.
		Mnemonic	: Specifies the index of underlined char and a key combination (Alt + char) which change the
						checkboxs's state and invokes the action listener. Default = -1.
		Alignment	: The alignment of the text and image similar to label. Possible values are: left, center, and
						right. Default = center.
		Selected	: The state of the checkbox. True if the checkbox is selected, false if it's not. Default = false.
		Group		: Identifies the radio button group if not null. Only one radio button at a time can be selected.
						User can set on a radio button, the selected button of the group will be set off (the group
						members is searched only in the same parent).
		Text		: The text string that the checkbox displays.
		GComponent	: @see GComponent.

	Listeners:
		Action		: Invokes the given method when the checkbox state is changed by mouse or keyboard event.

	Keyboard:
		Spacebar	: Selects or deselects.
		Tab			: Navigate forward.
		ShiftTab	: Navigate backward.
*/
type GCheckBox struct {
	Icon      p.Icon      `json:"icon"`
	Mnemonic  int         `json:"mnemonic"`
	Alignment p.Alignment `json:"alignment"`
	Selected  bool        `json:"selected"`
	Group     string      `json:"group"`
	Text      string      `json:"text"`
	w.GComponent
}

// DefaultCheckBox returns a check box structure with default values.
func DefaultCheckBox() *GCheckBox {
	return &GCheckBox{Selected: false, Alignment: p.ACenter, Mnemonic: -1}
}

// Action Invokes the given method when the checkbox state is changed by mouse or keyboard event.
func (gcb GCheckBox) Action() {
	fmt.Println(gcb)
}
