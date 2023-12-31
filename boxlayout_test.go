package fyneplus_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/must11/fyneplus"
	"github.com/stretchr/testify/assert"
)

// NewRectangle returns a new Rectangle instance
func NewMinSizeRect(min fyne.Size) *canvas.Rectangle {
	rect := &canvas.Rectangle{}
	rect.SetMinSize(min)

	return rect
}

func TestHBoxLayout_Simple(t *testing.T) {
	cellSize := fyne.NewSize(50, 50)

	obj1 := NewMinSizeRect(cellSize)
	obj2 := NewMinSizeRect(cellSize)
	obj3 := NewMinSizeRect(cellSize)
	layout := fyneplus.NewHBoxLayoutWithMargin(*fyneplus.NewMargin(100, 20, 20, 20))
	container := container.New(layout, obj1, obj2, obj3)

	// We are not in a window. Resize the container to a default size.
	t.Log(container.MinSize())
	container.Resize(container.MinSize())
	t.Log(container.Size())
	assert.Equal(t, container.MinSize(), fyne.NewSize(150+(theme.Padding()*2)+120, 50+40))

	assert.Equal(t, obj1.Size(), cellSize)
	cell2Pos := fyne.NewPos(50+theme.Padding(), 0)
	assert.Equal(t, cell2Pos, obj2.Position())
	cell3Pos := fyne.NewPos(100+theme.Padding()*2, 0)
	assert.Equal(t, cell3Pos, obj3.Position())
}
