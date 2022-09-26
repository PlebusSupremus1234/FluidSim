package particle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Particle) Integrate(dt float32) {
	// Integrate the particle's position
	p.X = rl.Vector2Add(p.X, rl.Vector2Scale(p.V, dt))
}
