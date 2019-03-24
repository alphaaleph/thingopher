// Package layouts handles all panel related widgets.
package layouts

/*
Desktop widget is the representation of a form's desktop.
	Parameters:

	Desktop is the root pane, it contains components (mainly panels, these fill the available space), dialogs (centered), combolists, popupmenus, and tooltips. You can add only components and dialogs to desktop, the last will be on the top. Desktop is also suitable to create MDI applications.
*/
type Desktop struct {
	component Component
}
