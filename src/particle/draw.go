package particle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func (p *Particle) Draw() {
	x := int32(math.Round(float64(p.X.X)))
	y := int32(math.Round(float64(p.X.Y)))

	rl.DrawRectangle(x-4, y-4, 8, 8, rl.Blue)
}
