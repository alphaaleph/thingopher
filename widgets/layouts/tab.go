// Package layouts handles all panel related widgets.
package layouts

/*
Tab widget is the representation of a form's tabbed pane.
	Parameters:
		mnemonic	: Specifies the underlined char and the key combination which selects the tab. Default = -1.

	TabbedPane contains tabs and components. The first tab (has an index equal to 0) is associated with the first component. This example tabbed pane has 3 tab and 3 components (text areas), tabs are on the left, and the second component is visible.

	<tabbedpane placement="left" selected="1" action="tabchanged">
  		<tab text="One" icon="image.gif">
    		<textarea text="One" />
  		</tab>
  		<tab text="Two" alignment="right">
    		<textarea text="Two" />
  		</tab>
  		<tab text="Three" enabled="false">
    		<textarea text="Three" />
  		</tab>
	</tabbedpane>
*/
type Tab struct {
	mnemonic int
	*Component
}

// KEYBOARD

// Action invokes the method when the tabbedpane has changed its selected tab index.
func (tp *TabbedPane) Action() {

}
