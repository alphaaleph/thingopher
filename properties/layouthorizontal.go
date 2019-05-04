// Package properties defines some of the types used by the widgets parameters.
package properties

import "strings"

//ChoiceHorizontal defines the various widget horizontal aligments.
type ChoiceHorizontal uint8

/*
	CHFill indicates that the component will use all the space.
	CHCenter indicates that the component will centered in the total space.
	CHLeft indicates that the component will be aligned on the left side.
	CHRight indicates that the component will be aligned on the right side.
*/
const (
	CHFill   ChoiceHorizontal = iota // 0x00
	CHCenter                         // 0x01
	CHLeft                           // 0x02
	CHRight                          // 0x03
)

//Text return the string representation of the horizontal choice.
func (ch ChoiceHorizontal) Text() string {
	switch ch {
	case CHFill:
		return "FILL"
	case CHCenter:
		return "CENTER"
	case CHLeft:
		return "LEFT"
	case CHRight:
		return "RIGHT"
	default:
		return ""
	}
}

//List returns a comma delimeted list of choice horizontal strings.
func (ch ChoiceHorizontal) List() string {
	list := ""

	//horizontal choices
	choiceHorizontal := map[ChoiceHorizontal]string{
		CHFill:   "FILL",
		CHCenter: "CENTER",
		CHLeft:   "LEFT",
		CHRight:  "RIGHT",
	}

	//set text
	for key, value := range choiceHorizontal {
		if (ch & key) != 0 {
			list += value + ","
		}
	}

	//remove the last comma and return the value
	list = strings.TrimSuffix(list, ",")
	return list
}
