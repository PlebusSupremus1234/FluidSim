package simulation

import (
	"fmt"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) drawScene(particles []*particle.Particle) {
	// Draw stats
	rl.DrawText(fmt.Sprintf("FPS: %.2f", rl.GetFPS()), 12, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Particles: %d", len(particles)), 12, 30, 20, rl.White)
}
