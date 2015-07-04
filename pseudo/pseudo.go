/**
 * pseudo color package
 */

package pseudo

import (
	"image/color"
	"math"
)

// ConvLinear RGB 線形関数
func ConvLinear(x int, y int, c color.Color) color.Color {
	r, g, b, _ := c.RGBA()

	l := float64(r)
	const dd = 0xFFFF / 4.0
	delta := dd

	if delta = dd * 3.0; delta < l {
		l -= delta
		r = 0xFFFF
		g = 0xFFFF - uint32(l*4.0)
		b = 0
	} else if delta = dd * 2.0; delta < l {
		l -= delta
		r = uint32(l * 4.0)
		g = 0xFFFF
		b = 0
	} else if delta = dd; delta < l {
		l -= delta
		r = 0
		g = 0xFFFF
		b = 0xFFFF - uint32(l*4.0)
	} else {
		r = 0
		g = uint32(l * 4.0)
		b = 0xFFFF
	}

	return color.RGBA64{uint16(r), uint16(g), uint16(b), 0xFFFF}
}

// ConvSigmoid RGB 非線形関数 sigmoid
func ConvSigmoid(_ int, _ int, c color.Color) color.Color {
	r, g, b, _ := c.RGBA()

	l := float64(r)
	const dd = 0xFFFF / 4
	const dr = 0.0002

	r = uint32(float64(0xFFFF) * 1 / (1 + math.Exp(dr*(-l+2*dd+dd/2))))
	b = uint32(float64(0xFFFF) * (1 - 1/(1+math.Exp(dr*(-l+dd+dd/2)))))

	if l < dd*2 {
		g = uint32(float64(0xFFFF) * 1 / (1 + math.Exp(dr*(-l+dd/2))))
	} else {
		g = uint32(float64(0xFFFF) * (1 - 1/(1+math.Exp(dr*(-l+3*dd+dd/2)))))
	}

	return color.RGBA64{uint16(r), uint16(g), uint16(b), 0xFFFF}
}

// ConvSin RGB 非線形関数 sin
func ConvSin(_ int, _ int, c color.Color) color.Color {
	r, g, b, _ := c.RGBA()

	l := float64(r) / 0xFFFF
	const shift = math.Pi + math.Pi/4

	r = uint32(0xFFFF * (math.Sin(1.5*math.Pi*l+shift) + 1.0) / 2.0)
	g = uint32(0xFFFF * (math.Sin(1.5*math.Pi*l+shift+math.Pi/2.0) + 1.0) / 2.0)
	b = uint32(0xFFFF * (math.Sin(1.5*math.Pi*l+shift+math.Pi) + 1.0) / 2.0)

	return color.RGBA64{uint16(r), uint16(g), uint16(b), 0xFFFF}
}
