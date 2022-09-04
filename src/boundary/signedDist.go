package boundary

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

// SignedDistance function
func SignedDistance(x rl.Vector2, bounds []*Boundary) (rl.Vector2, float32) {
	// Find the closest point on the rectangle to the particle using vector projection
	dist := float32(math.Inf(1))
	closest := rl.Vector2Zero()

	inBoundary := false

	// Loop through all boundaries
	for _, i := range bounds {
		if i.Contains(x) {
			inBoundary = true
		}

		// Loop through sides of the boundary
		for j := 0; j < len(i.vertices); j++ {
			c := closestPoint(x, i.vertices[j], i.vertices[(j+1)%len(i.vertices)])

			d := rl.Vector2Distance(x, c)

			if d < dist {
				dist = d
				closest = c
			}
		}
	}

	if inBoundary {
		return closest, -dist
	} else {
		return closest, dist
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
