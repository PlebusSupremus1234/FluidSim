package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	"math/rand"

	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//func initParticles(H, cols, rows float32) ([]*particle.Particle, int) {
//	var particles []*particle.Particle
//
//	// Initial particle configuration
//	for i := 0; i < 30; i++ {
//		for j := 0; j < 20; j++ {
//			particles = append(particles, particle.New(
//				100+H*float32(j)+rand.Float32(),
//				200+H*float32(i)+rand.Float32(),
//
//				len(particles), // Index
//			))
//		}
//	}
//
//	return particles, len(particles) - 1
//}

func initBoundaries(width, height, H float32) []*boundary.Line {
	var boundaries []*boundary.Line

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{},
		rl.Vector2{Y: height},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{Y: height},
		rl.Vector2{X: width, Y: height},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: width, Y: height},
		rl.Vector2{X: width},
		H,
	))

	boundaries = append(boundaries, boundary.New(
		rl.Vector2{X: width},
		rl.Vector2{},
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

			x := mX + float32(j)*s.H + rand.Float32()
			y := mY + float32(i)*s.H + rand.Float32()

			badX := (x+s.H > s.ViewWidth) || (x-s.H < 0)
			badY := (y+s.H > s.ViewHeight) || (y-s.H < 0)

			if !badX && !badY {
				newParticle := particle.New(x, y, s.LatestIndex+1)

				s.LatestIndex++
				s.particles = append(s.particles, newParticle)
			}
		}
	}
}
