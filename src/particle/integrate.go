package particle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Particle) Integrate(dt float32) {
	// Integrate the laws of motion
	p.V = rl.Vector2Add(p.V, rl.Vector2Scale(p.A, dt/p.Rho))
	p.X = rl.Vector2Add(p.X, rl.Vector2Scale(p.V, dt))
}
