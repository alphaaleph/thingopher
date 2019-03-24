// Package display manage all non-input widgets.
package display

/*
ProgressBar widget is the representation of a form's progress bar.
	Parameters:
		orientation	: Possible values are: horizontal, and vertical. Default = horizontal.
		minimum		: The progressbar's minimum value. Default = 0.
		maximum		: The progressbar's maximum value. Default = 100.
		value		: The value is always between the progressbar's minimum and maximum values, inclusive. By default, the value equals the minimum. Default = 0.

	ProressBar displays an integer value within a bounded interval, the following between 25 and 75. Its value is currently in the middle.

	<progressbar minimum="25" maximum="75" value="50" orientation="vertical" />
*/
type ProgressBar struct {
	orientation Choice
	minimum     int
	maximum     int
	value       int
	component   Component
}
