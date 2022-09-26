package particle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	X rl.Vector2 // Position
	V rl.Vector2 // Velocity

	Rho float32 // Density
	P   float32 // Pressure
	M   float32 // Mass

	Index int // Index
}

func New(x, y float32, index int) *Particle {
	return &Particle{
		X: rl.NewVector2(x, y),
		V: rl.Vector2Zero(),

		Rho: 0,
		P:   0,
		M:   3,

		Index: index,
	}
}
