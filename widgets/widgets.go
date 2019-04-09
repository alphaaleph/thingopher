// Package widgets holds all the components supported by Thinggopher
package widgets

/*
Widget interface provides the basic widgets calls.
	Calls:
		Init		: A method to invoke only once when the loading of the xml resource is finished.
		FocusLost	: Invoked when a widget loses the keyboard focus, thusd is no longer the focus owner.
		FocusGained	: Invoked when a widget gains the keyboard focus, thus is now the focus owner.
*/
type Widget interface {
	Init()
	FocusLost()
	FocusGained()
}
