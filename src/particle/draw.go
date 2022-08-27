package particle

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Particle) Draw() {
	x := int32(math.Round(float64(p.X.X)))
	y := int32(math.Round(float64(p.X.Y)))

	color := rl.Blue
	if p.T == Bound {
		color = rl.Gray
	}

	rl.DrawRectangle(x-4, y-4, 8, 8, color)
}
