package simulation

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) drawScene(count int) {
	// Draw simulation stats
	rl.DrawText(fmt.Sprintf("FPS: %.2f", rl.GetFPS()), 12, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Particles: %d", count), 12, 30, 20, rl.White)
}
