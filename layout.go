package fyneplus

import "fyne.io/fyne/v2"

type Layout interface {
	fyne.Layout
	Margin(m margin)
	GetMargin() margin
}
