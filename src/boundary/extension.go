package boundary

func ExtensionFunc(x, h, f float32) float32 {
	if 0 < x && x < h {
		return cubicSpline(x, h, f) / cubicSpline(0, h, f)
	} else if x <= 0 {
		return 1
	} else {
		return 0
	}
}
