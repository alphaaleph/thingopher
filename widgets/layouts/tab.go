// Package layouts handles all panel related widgets.
package layouts

/*
Tab widget is the representation of a form's tabbed pane.
	Parameters:
		mnemonic	: Specifies the underlined char and the key combination which selects the tab. Default = -1.

	Tabs are identified in the component's list as 'tab', and components with the 'component' key.
*/
type Tab struct {
	mnemonic  int
	component ComboBox
}
