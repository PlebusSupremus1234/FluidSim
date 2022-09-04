package boundary

func cubicSpline(r, h, f float32) float32 {
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
