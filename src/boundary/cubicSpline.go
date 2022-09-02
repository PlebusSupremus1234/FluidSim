package boundary

import "math"

func cubicSpline(r, h float32) float32 {
	f := 40 / (7 * math.Pi * h * h)
	q := r / h

	if 0 <= q && q <= 0.5 {
		return f * (6*(q*q*q-q*q) + 1)
	} else if 0.5 < q && q <= 1 {
		a := 1 - q
		return f * 2 * a * a * a
	} else {
		return 0
	}
}
