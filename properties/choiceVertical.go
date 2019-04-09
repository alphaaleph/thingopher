// Package properties defines some of the types used by the widgets parameters.
package properties

import "strings"

//ChoiceVertical defines the various widget vertical aligments.
type ChoiceVertical uint8

/*
	CVFill indicates that the component will use all the space.
	CVCenter indicates that the component will centered in the total space.
	CVTop indicates that the component will be aligned on the top side.
	CVBottom indicates that the component will be aligned on the bottom side.
*/
const (
	CVFill   ChoiceVertical = iota // 0x00
	CVCenter                       // 0x01
	CVTop                          // 0x02
	CVBottom                       // 0x03
)

//Text return the string representation of the vertical choice.
func (ch ChoiceVertical) Text() string {
	switch ch {
	case CVFill:
		return "FILL"
	case CVCenter:
		return "CENTER"
	case CVTop:
		return "TOP"
	case CVBottom:
		return "BOTTOM"
	default:
		return ""
	}
}

//List returns a comma delimeted list of choice vertical strings.
func (ch ChoiceVertical) List() string {
	list := ""

	//vertical choices
	choiceVertical := map[ChoiceVertical]string{
		CVFill:   "FILL",
		CVCenter: "CENTER",
		CVTop:    "TOP",
		CVBottom: "BOTTOM",
	}

	//set text
	for key, value := range choiceVertical {
		if (ch & key) != 0 {
			list += value + ","
		}
	}

	//remove the last comma and return the value
	list = strings.TrimSuffix(list, ",")
	return list
}
