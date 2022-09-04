package renderer

import (
	"fmt"
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func Draw(particles []*particle.Particle, boundaries []*boundary.Boundary) {
	// Draw particles
	for _, p := range particles {
		x := int32(math.Round(float64(p.X.X)))
		y := int32(math.Round(float64(p.X.Y)))

		rl.DrawRectangle(x-4, y-4, 8, 8, rl.Blue)
		//rl.DrawCircleV(p.X, 16, rl.Blue)
	}

	for _, b := range boundaries {
		b.Draw()
	}

	// Stats
	rl.DrawText(fmt.Sprintf("FPS: %.2f", rl.GetFPS()), 12, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Particles: %d", len(particles)), 12, 30, 20, rl.White)
}
