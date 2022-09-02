package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) integrate() {
	for _, i := range s.particles {
		i.V = rl.Vector2Add(i.V, rl.Vector2Scale(i.A, s.Dt/i.Rho))
		i.X = rl.Vector2Add(i.X, rl.Vector2Scale(i.V, s.Dt))
	}
}
