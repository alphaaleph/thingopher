// Package properties defines some of the types used by the widgets parameters.
package properties

// GPlacement defines the various widgets placements.
type GPlacement uint8

/*
	Top indicates that the menu bar will be placed at the top of a container.
	Bottom indicates that the menu bar will be placed at the bottom of a container.
*/
const (
	Top    GPlacement = iota // 0x00
	Bottom                   // 0x01
)

//Text returns the string representation of the placement choice.
func (gpl GPlacement) Text() string {
	switch gpl {
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	default:
		return ""
	}
}

//List returns a slice of placement strings.
func (gpl GPlacement) List() []string {
	return []string{Top.Text(), Bottom.Text()}
}
