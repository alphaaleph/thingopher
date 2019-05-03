// Package properties defines some of the types used by the widgets parameters.
package properties

//GSelection defines the various widget selections.
type GSelection uint8

/*
	Single
	Interval
	Multiple
	ACenter indicates that the text will be centered in the widget.
	ALeft indicates that the text will be aligned on the left side of the widget.
	ARight indicates that the text will be aligned on the right side of the widget.
*/
const (
	Single   GSelection = iota // 0x00
	Interval                   // 0x01
	Multiple                   // 0x02
)

//Text returns the string representation of the selection choice.
func (gal GSelection) Text() string {
	switch gal {
	case Single:
		return "Single"
	case Interval:
		return "Interval"
	case Multiple:
		return "Multiple"
	default:
		return ""
	}
}

//List returns a slice of selection strings.
func (gal GSelection) List() []string {
	return []string{Single.Text(), Interval.Text(), Multiple.Text()}
}
