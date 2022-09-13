package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

func initBoundaries(width, height, H float32) []*boundary.Boundary {
	var boundaries []*boundary.Boundary

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: H / 2},
		rl.Vector2{X: H / 2, Y: height - H},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{Y: height - H/2},
		rl.Vector2{X: width - H, Y: height - H/2},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: width - H/2, Y: height},
		rl.Vector2{X: width - H/2, Y: H},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: width, Y: H / 2},
		rl.Vector2{X: H, Y: H / 2},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: 200, Y: 200},
		rl.Vector2{X: 500, Y: 500},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: 600, Y: 500},
		rl.Vector2{X: 1000, Y: 500},
		H,
	))

	return boundaries
}

func (s *Simulation) SpawnParticles() {
	for i := -5; i <= 5; i++ {
		for j := -5; j <= 5; j++ {
			mousePos := rl.GetMousePosition()
			mX, mY := mousePos.X, mousePos.Y

			x := mX + float32(j)*s.h + rand.Float32()
			y := mY + float32(i)*s.h + rand.Float32()

			badX := (x+s.h > s.viewW) || (x-s.h < 0)
			badY := (y+s.h > s.viewH) || (y-s.h < 0)

			if !badX && !badY {
				newParticle := particle.New(x, y, s.index+1)

				s.index++
				s.particles = append(s.particles, newParticle)
			}
		}
	}

	//mousePos := rl.GetMousePosition()
	//mX, mY := mousePos.X, mousePos.Y
	//
	//x := mX
	//y := mY
	//
	//badX := (x+s.h > s.viewW) || (x-s.h < 0)
	//badY := (y+s.h > s.viewH) || (y-s.h < 0)
	//
	//if !badX && !badY {
	//	newParticle := particle.New(x, y, s.index+1)
	//
	//	s.index++
	//	s.particles = append(s.particles, newParticle)
	//}
}
