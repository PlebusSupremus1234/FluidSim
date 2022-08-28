package renderer

import (
	"fmt"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func Draw(particles []*particle.Particle) {
	// Draw particles
	for _, p := range particles {
		x := int32(math.Round(float64(p.X.X)))
		y := int32(math.Round(float64(p.X.Y)))

		color := rl.Blue
		if p.T == particle.Bound {
			color = rl.Gray
		}

		rl.DrawRectangle(x-4, y-4, 8, 8, color)
	}

	// Stats
	rl.DrawText(fmt.Sprintf("FPS: %.2f", rl.GetFPS()), 12, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Particles: %d", len(particles)), 12, 30, 20, rl.White)
}
