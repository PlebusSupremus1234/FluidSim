package boundary

func ExtensionFunc(x, h float32) float32 {
	if 0 < x && x < h {
		return cubicSpline(x, h) / cubicSpline(0, h)
	} else if x <= 0 {
		return 1
	} else {
		return 0
	}
}
