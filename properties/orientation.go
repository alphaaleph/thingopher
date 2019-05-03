// Package properties defines some of the types used by the widgets parameters.
package properties

//GOrientation defines the various widgets orientations.
type GOrientation uint8

/*
	Horizontal indicates that a widget will be displayed in horizontal fashion.
	Vertical indicates that a widget will be displayed in vertical fashion.
*/
const (
	Horizontal GOrientation = iota // 0x00
	Vertical                       // 0x01
)

//Text returns the string representation of the orientation choice.
func (gor GOrientation) Text() string {
	switch gor {
	case Horizontal:
		return "Horizontal"
	case Vertical:
		return "Vertical"
	default:
		return ""
	}
}

//List returns a slice of orientation strings.
func (gor GOrientation) List() []string {
	return []string{Horizontal.Text(), Vertical.Text()}
}
