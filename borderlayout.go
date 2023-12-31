package fyneplus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Declare conformity with Layout interface
var _ Layout = (*borderLayout)(nil)

type borderLayout struct {
	top, bottom, left, right fyne.CanvasObject
	margin                   margin
}

func (l *borderLayout) Margin(m margin) {
	l.margin.Update(m.Size())
}

func (l *borderLayout) GetMargin() margin {
	return l.margin
}

// NewBorderLayout creates a new BorderLayout instance with top, bottom, left
// and right objects set. All other items in the container will fill the centre
// space
func NewBorderLayout(top, bottom, left, right fyne.CanvasObject) fyne.Layout {
	return &borderLayout{top, bottom, left, right, *NewMargin(0, 0, 0, 0)}
}
func NewBorderLayoutWithMargin(top, bottom, left, right fyne.CanvasObject, m margin) fyne.Layout {
	return &borderLayout{top, bottom, left, right, *NewMargin(m.Size())}
}

// Layout is called to pack all child objects into a specified size.
// For BorderLayout this arranges the top, bottom, left and right widgets at
// the sides and any remaining widgets are maximised in the middle space.
func (b *borderLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	padding := theme.Padding()
	var topSize, bottomSize, leftSize, rightSize fyne.Size
	if b.top != nil && b.top.Visible() {
		topHeight := b.top.MinSize().Height
		b.top.Resize(fyne.NewSize(size.Width, topHeight))
		b.top.Move(fyne.NewPos(b.margin.left, b.margin.top))
		topSize = fyne.NewSize(size.Width-b.margin.left, topHeight+padding)
	}
	if b.bottom != nil && b.bottom.Visible() {
		bottomHeight := b.bottom.MinSize().Height
		b.bottom.Resize(fyne.NewSize(size.Width, bottomHeight))
		b.bottom.Move(fyne.NewPos(b.margin.left, size.Height-bottomHeight-b.margin.bottom))
		bottomSize = fyne.NewSize(size.Width-b.margin.left, bottomHeight+padding)
	}
	if b.left != nil && b.left.Visible() {
		leftWidth := b.left.MinSize().Width
		b.left.Resize(fyne.NewSize(leftWidth, size.Height-topSize.Height-bottomSize.Height-b.margin.top-b.margin.bottom))
		b.left.Move(fyne.NewPos(b.margin.left, b.margin.top+topSize.Height))
		leftSize = fyne.NewSize(leftWidth+padding, size.Height-topSize.Height-bottomSize.Height-b.margin.top-b.margin.bottom)
	}
	if b.right != nil && b.right.Visible() {
		rightWidth := b.right.MinSize().Width
		b.right.Resize(fyne.NewSize(rightWidth, size.Height-topSize.Height-bottomSize.Height-b.margin.top-b.margin.bottom))
		b.right.Move(fyne.NewPos(size.Width-rightWidth-b.margin.right, topSize.Height+b.margin.top))
		rightSize = fyne.NewSize(rightWidth+padding, size.Height-topSize.Height-bottomSize.Height-b.margin.top-b.margin.bottom)
	}

	middleSize := fyne.NewSize(size.Width-leftSize.Width-rightSize.Width-b.margin.left-b.margin.right, size.Height-topSize.Height-bottomSize.Height-b.margin.top-b.margin.bottom)
	middlePos := fyne.NewPos(leftSize.Width+b.margin.left, topSize.Height+b.margin.top)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if child != b.top && child != b.bottom && child != b.left && child != b.right {
			child.Resize(middleSize)
			child.Move(middlePos)
		}
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For BorderLayout this is determined by the MinSize height of the top and
// plus the MinSize width of the left and right, plus any padding needed.
// This is then added to the union of the MinSize for any remaining content.
func (b *borderLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if child != b.top && child != b.bottom && child != b.left && child != b.right {
			minSize = minSize.Max(child.MinSize())
		}
	}

	padding := theme.Padding()

	if b.left != nil && b.left.Visible() {
		leftMin := b.left.MinSize()
		minHeight := fyne.Max(minSize.Height, leftMin.Height)
		minSize = fyne.NewSize(minSize.Width+leftMin.Width+padding, minHeight)
	}
	if b.right != nil && b.right.Visible() {
		rightMin := b.right.MinSize()
		minHeight := fyne.Max(minSize.Height, rightMin.Height)
		minSize = fyne.NewSize(minSize.Width+rightMin.Width+padding, minHeight)
	}

	if b.top != nil && b.top.Visible() {
		topMin := b.top.MinSize()
		minWidth := fyne.Max(minSize.Width, topMin.Width)
		minSize = fyne.NewSize(minWidth, minSize.Height+topMin.Height+padding)
	}
	if b.bottom != nil && b.bottom.Visible() {
		bottomMin := b.bottom.MinSize()
		minWidth := fyne.Max(minSize.Width, bottomMin.Width)
		minSize = fyne.NewSize(minWidth, minSize.Height+bottomMin.Height+padding)
	}
	minSize.Width = minSize.Width + b.margin.left + b.margin.right
	minSize.Height = minSize.Height + b.margin.top + b.margin.bottom
	return minSize
}
