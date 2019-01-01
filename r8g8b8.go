package texture

import (
	"image"
	"image/color"

	"github.com/cozely/colour"
)

////////////////////////////////////////////////////////////////////////////////

type R8G8B8 struct {
	Pix    []uint8
	Stride int
	Rect   image.Rectangle
}

func NewR8G8B8(r image.Rectangle) *R8G8B8 {
	w, h := r.Dx(), r.Dy()
	return &R8G8B8{
		Pix:    make([]uint8, 3*w*h, 3*w*h),
		Stride: 3 * w,
		Rect:   r,
	}
}

func (t *R8G8B8) ColorModel() color.Model {
	return colour.R8G8B8Model
}

func (t *R8G8B8) Bounds() image.Rectangle {
	return t.Rect
}

func (t *R8G8B8) At(x, y int) color.Color {
	return t.R8G8B8At(x, y)
}

func (t *R8G8B8) R8G8B8At(x, y int) colour.R8G8B8 {
	i := t.PixOffset(x, y)
	return colour.R8G8B8{
		R: t.Pix[i+0],
		G: t.Pix[i+1],
		B: t.Pix[i+2],
	}
}

func (t *R8G8B8) SetRGBAt(x, y int, c colour.R8G8B8) {
	cc := colour.R8G8B8Model.Convert(c).(colour.R8G8B8)
	t.SetR8G8B8At(x, y, cc)
}

func (t *R8G8B8) SetR8G8B8At(x, y int, c colour.R8G8B8) {
	i := t.PixOffset(x, y)
	t.Pix[i+0] = c.R
	t.Pix[i+1] = c.G
	t.Pix[i+2] = c.B
}

func (t *R8G8B8) PixOffset(x, y int) int {
	return (y-t.Rect.Min.Y)*t.Stride + (x-t.Rect.Min.X)*3
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (t *R8G8B8) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(t.Rect)
	// Needed to avoid index out of range (in Pix below):
	if r.Empty() {
		return &R8G8B8{}
	}
	i := t.PixOffset(r.Min.X, r.Min.Y)
	return &R8G8B8{
		Pix:    t.Pix[i:],
		Stride: t.Stride,
		Rect:   r,
	}
}

func (t *R8G8B8) Opaque() bool {
	return true
}
