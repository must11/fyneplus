package fyneplus

import (
	"fyne.io/fyne/v2"
	ly "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

// Declare conformity with Layout interface
var _ Layout = (*boxLayout)(nil)

type boxLayout struct {
	horizontal bool
	margin     margin
}

func (l *boxLayout) GetMargin() margin {
	return l.margin
}

func (l *boxLayout) Margin(m margin) {
	l.margin.Update(m.Size())
}

// NewHBoxLayout returns a horizontal box layout for stacking a number of child
// canvas objects or widgets left to right. The objects are always displayed
// at their horizontal MinSize. Use a different layout if the objects are intended
// to be larger then their horizontal MinSize.
func NewHBoxLayout() Layout {
	return &boxLayout{true, *NewMargin(0, 0, 0, 0)}
}
func NewHBoxLayoutWithMargin(m margin) Layout {
	return &boxLayout{true, *NewMargin(m.Size())}
}

func NewVBoxLayoutWithMargin(m margin) Layout {
	return &boxLayout{false, *NewMargin(m.Size())}
}

// NewVBoxLayout returns a vertical box layout for stacking a number of child
// canvas objects or widgets top to bottom. The objects are always displayed
// at their vertical MinSize. Use a different layout if the objects are intended
// to be larger then their vertical MinSize.
func NewVBoxLayout() fyne.Layout {
	return &boxLayout{false, *NewMargin(0, 0, 0, 0)}
}

func (g *boxLayout) isSpacer(obj fyne.CanvasObject) bool {
	if !obj.Visible() {
		return false // invisible spacers don't impact layout
	}

	spacer, ok := obj.(ly.SpacerObject)
	if !ok {
		return false
	}

	if g.horizontal {
		return spacer.ExpandHorizontal()
	}

	return spacer.ExpandVertical()
}

// Layout is called to pack all child objects into a specified size.
// For a VBoxLayout this will pack objects into a single column where each item
// is full width but the height is the minimum required.
// Any spacers added will pad the view, sharing the space if there are two or more.
func (g *boxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	spacers := 0
	visibleObjects := 0
	// Size taken up by visible objects
	total := float32(0)

	for _, child := range objects {
		if !child.Visible() {
			continue
		}
		if g.isSpacer(child) {
			spacers++
			continue
		}

		visibleObjects++
		if g.horizontal {
			total += child.MinSize().Width
		} else {
			total += child.MinSize().Height
		}
	}

	padding := theme.Padding()

	// Amount of space not taken up by visible objects and inter-object padding
	var extra float32
	if g.horizontal {
		extra = size.Width - total - (padding * float32(visibleObjects-1)) - g.margin.left - g.margin.right
	} else {
		extra = size.Height - total - (padding * float32(visibleObjects-1)) - g.margin.top - g.margin.bottom
	}

	// Spacers split extra space equally
	spacerSize := float32(0)
	if spacers > 0 {
		spacerSize = extra / float32(spacers)
	}

	x, y := g.margin.left, g.margin.top
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if g.isSpacer(child) {
			if g.horizontal {
				x += spacerSize
			} else {
				y += spacerSize
			}
			continue
		}
		child.Move(fyne.NewPos(x, y))

		if g.horizontal {
			width := child.MinSize().Width
			x += padding + width
			child.Resize(fyne.NewSize(width, size.Height))
		} else {
			height := child.MinSize().Height
			y += padding + height
			child.Resize(fyne.NewSize(size.Width, height))
		}
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For a BoxLayout this is the width of the widest item and the height is
// the sum of of all children combined with padding between each.
func (g *boxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	addPadding := false
	padding := theme.Padding()
	for _, child := range objects {
		if !child.Visible() || g.isSpacer(child) {
			continue
		}

		childMin := child.MinSize()
		if g.horizontal {
			minSize.Height = fyne.Max(childMin.Height, minSize.Height)
			minSize.Width += childMin.Width
			if addPadding {
				minSize.Width += padding
			}
		} else {
			minSize.Width = fyne.Max(childMin.Width, minSize.Width)
			minSize.Height += childMin.Height
			if addPadding {
				minSize.Height += padding
			}
		}
		addPadding = true
	}
	minSize.Width = minSize.Width + g.margin.left + g.margin.right
	minSize.Height = minSize.Height + g.margin.top + g.margin.bottom
	return minSize
}
