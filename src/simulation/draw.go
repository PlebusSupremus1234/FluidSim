package simulation

import (
	"fmt"
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) Draw(particles []*particle.Particle, boundaries []*boundary.Boundary) {
	// Draw boundaries
	for _, b := range boundaries {
		b.Draw()
	}

	// Stats
	rl.DrawText(fmt.Sprintf("FPS: %.2f", rl.GetFPS()), 12, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Particles: %d", len(particles)), 12, 30, 20, rl.White)
}
