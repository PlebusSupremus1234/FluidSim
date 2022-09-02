package boundary

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Line struct {
	A rl.Vector2
	B rl.Vector2

	vertices []rl.Vector2
}

func New(a, b rl.Vector2, w float32) *Line {
	// Boundaries are rectangles but defined by two points
	c := math.Atan2(float64(b.Y-a.Y), float64(b.X-a.X)) + math.Pi/2 // Angle between the two points
	d := c - math.Pi/2                                              // Rotation angle for each vertex

	cornerA := rl.NewVector2(a.X-w/2, a.Y-w/2)
	cornerB := rl.NewVector2(a.X-w/2, a.Y+w/2)
	cornerC := rl.NewVector2(b.X+w/2, b.Y+w/2)
	cornerD := rl.NewVector2(b.X+w/2, b.Y-w/2)

	rotatedA := rotatePoint(cornerA, a, d)
	rotatedB := rotatePoint(cornerB, a, d)
	rotatedC := rotatePoint(cornerC, b, d)
	rotatedD := rotatePoint(cornerD, b, d)

	return &Line{
		A: a,
		B: b,

		vertices: []rl.Vector2{
			rotatedA,
			rotatedB,
			rotatedC,
			rotatedD,
		},
	}
}

func rotatePoint(pos, origin rl.Vector2, angle float64) rl.Vector2 {
	c := float32(math.Cos(angle))
	s := float32(math.Sin(angle))

	x := pos.X - origin.X
	y := pos.Y - origin.Y

	xNew := x*c - y*s
	yNew := x*s + y*c

	return rl.NewVector2(xNew+origin.X, yNew+origin.Y)
}

func (l *Line) Draw() {
	rl.DrawLineEx(l.A, l.B, 16, rl.Red)
}

func (l *Line) Contains(x rl.Vector2) bool {
	collision := false
	next := 0

	for current := 0; current < len(l.vertices); current++ {
		next = current + 1
		if next == len(l.vertices) {
			next = 0
		}

		vc := l.vertices[current]
		vn := l.vertices[next]

		if ((vc.Y >= x.Y && vn.Y < x.Y) || (vc.Y < x.Y && vn.Y >= x.Y)) &&
			(x.X < (vn.X-vc.X)*(x.Y-vc.Y)/(vn.Y-vc.Y)+vc.X) {
			collision = !collision
		}
	}
	return collision
}
