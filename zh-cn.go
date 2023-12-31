package fyneplus

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type simpleChineseTheme struct{}

var _ fyne.Theme = (*simpleChineseTheme)(nil)

func (t *simpleChineseTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(c, v)
}
func (t *simpleChineseTheme) Font(ts fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(ts)
}
func (t *simpleChineseTheme) Icon(icon fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(icon)
}
func (t *simpleChineseTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}
