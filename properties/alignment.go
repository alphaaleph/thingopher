// Package properties defines some of the types used by the widgets parameters.
package properties

//GAlignment defines the various widget horizontal aligments.
type GAlignment uint8

/*
	ACenter indicates that the text will be centered in the widget.
	ALeft indicates that the text will be aligned on the left side of the widget.
	ARight indicates that the text will be aligned on the right side of the widget.
*/
const (
	ACenter GAlignment = iota // 0x00
	ALeft                     // 0x01
	ARight                    // 0x02
)

//Text returns the string representation of the alignment choice.
func (gal GAlignment) Text() string {
	switch gal {
	case ACenter:
		return "Center"
	case ALeft:
		return "Left"
	case ARight:
		return "Right"
	default:
		return ""
	}
}

//List returns a slice of alignment strings.
func (gal GAlignment) List() []string {
	return []string{ACenter.Text(), ALeft.Text(), ARight.Text()}
}
