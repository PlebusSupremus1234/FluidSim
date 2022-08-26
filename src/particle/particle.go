package particle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PARTICLE = iota
	BOUNDARY
)

type Particle struct {
	X rl.Vector2 // Position
	V rl.Vector2 // Velocity
	A rl.Vector2 // Acceleration

	Rho float32 // Density
	P   float32 // Pressure
	M   float32 // Mass

	Index int // Grid Index
	T     int // Type
}

func New(x, y float32, index, T int) *Particle {
	return &Particle{
		X: rl.NewVector2(x, y),
		V: rl.Vector2Zero(),
		A: rl.Vector2Zero(),

		Rho: 0,
		P:   0,
		M:   2.5,

		Index: index,
		T:     T,
	}
}
