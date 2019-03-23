// Package layouts handles all panel related widgets.
package layouts

/*
TabbedPane widget is the representation of a form's tabbed pane.
	Parameters:
		placement 	: The placement for the tabs relative to the content. Possible values are: top, left, bottom, right, and stack. The stack placement arranges tabs so that they resemble a sidebar. Default = top.
		selected	: The index of the currently selected tab, and the visible content (the first index is 0). Value -1 means there is no selected tab.

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
type TabbedPane struct {
	placement Choice
	selected  int
	*Component
}

// LISTENERS

// Action invokes the method when the tabbedpane has changed its selected tab index.
func (tp *TabbedPane) Action() {

}
