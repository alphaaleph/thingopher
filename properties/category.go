// Package properties defines some of the types used by the widgets parameters.
package properties

//Category defines the various widget's control and look.
type Category uint8

/*
	CNormal is the default value of a widget's category.
	CDefault is used as a widget's category control.
	CCancel is used as a widget's category control.
	CLink changes the appereance of a button so that it resembles HTML link.
*/
const (
	CNormal  Category = iota // 0x00
	CDefault                 // 0x01
	CCancel                  // 0x02
	CLink                    // 0x03
)

//Text returns the string representation of the widget's category.
func (cat Category) Text() string {
	switch cat {
	case CNormal:
		return "Normal"
	case CDefault:
		return "Default"
	case CCancel:
		return "Cancel"
	case CLink:
		return "Link"
	default:
		return ""
	}
}

//List returns a slice of category strings.
func (cat Category) List() []string {
	return []string{CNormal.Text(), CDefault.Text(), CCancel.Text(), CLink.Text()}
}
