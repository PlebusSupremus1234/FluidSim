package particle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func (p *Particle) Draw(scale float32) {
	// Draw the particle
	x := int32(math.Round(float64(p.X.X * scale)))
	y := int32(math.Round(float64(p.X.Y * scale)))

	rl.DrawRectangle(x-4, y-4, 8, 8, rl.Blue)
}
