package boundary

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

// Sdf is the signed distance function
func Sdf(x rl.Vector2, bounds []*Line) (rl.Vector2, float32) {
	// Find the closest point on the rectangle to the particle using vector projection
	dist := float32(math.Inf(1))
	closest := rl.Vector2Zero()

	inBoundary := false

	// Loop through all boundary lines
	for _, i := range bounds {
		if i.Contains(x) {
			inBoundary = true
		}

		a := i.A
		b := i.B
		c := closestPoint(x, a, b)
		d := rl.Vector2Subtract(x, c)

		if rl.Vector2LenSqr(d) < dist {
			dist = rl.Vector2LenSqr(d)
			closest = c
		}
	}

	infimum := float32(math.Sqrt(float64(dist)))

	if inBoundary {
		return closest, -infimum
	} else {
		return closest, infimum
	}
}

// closestPoint is a function to find the closest point on a line to a point
func closestPoint(p, a, b rl.Vector2) rl.Vector2 {
	w := rl.Vector2Subtract(a, b)
	v := rl.Vector2Subtract(p, b)
	c := rl.Vector2DotProduct(w, v)

	t := c / rl.Vector2LenSqr(w)
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}

	return rl.Vector2Add(b, rl.Vector2Scale(w, t))
}
