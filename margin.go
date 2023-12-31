package fyneplus

type margin struct {
	left   float32
	top    float32
	right  float32
	bottom float32
}

func NewMargin(left, top, right, bottom float32) *margin {
	return &margin{left: left, top: top, right: right, bottom: bottom}
}
func (m *margin) Size() (left, top, right, bottom float32) {
	return m.left, m.top, m.right, m.bottom
}
func (m *margin) Update(left, top, right, bottom float32) {
	m.left = left
	m.top = top
	m.right = right
	m.bottom = bottom
}
