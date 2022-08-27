package simulation

import (
	"math/rand"

	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func initParticles(H, cols, rows float32) []*particle.Particle {
	var particles []*particle.Particle

	// Initial particle configuration
	for i := 0; i < 30; i++ {
		for j := 0; j < 20; j++ {
			particles = append(particles, particle.New(
				100+H*float32(j)+rand.Float32(),
				200+H*float32(i)+rand.Float32(),

				len(particles), // Index

				particle.PARTICLE, // Type
			))
		}
	}

	// Outer boundary
	// Top and bottom
	for i := 0; i < 2; i++ {
		for j := 0; float32(j) < 2*cols; j++ {
			x := float32(j)*H/2 + H/4
			y := H/4 + float32(i)*(rows*H-H/2)

			particles = append(particles, particle.New(
				x, y,

				len(particles),

				particle.BOUNDARY,
			))
		}
	}

	// Left and right
	for i := 0; float32(i) < 2*rows; i++ {
		for j := 0; j < 2; j++ {
			x := H/4 + float32(j)*(cols*H-H/2)
			y := H/4 + float32(i)*H/2
			particles = append(particles, particle.New(
				x, y,

				len(particles),
				particle.BOUNDARY,
			))
		}
	}

	return particles
}

func (s *Simulation) SpawnParticles() {
	for i := -5; i <= 5; i++ {
		for j := -5; j <= 5; j++ {
			mousePos := rl.GetMousePosition()
			mX, mY := mousePos.X, mousePos.Y

			x := mX + float32(j)*s.H + rand.Float32()
			y := mY + float32(i)*s.H + rand.Float32()

			badX := (x+s.Eps > s.ViewWidth) || (x-s.Eps < 0)
			badY := (y+s.Eps > s.ViewHeight) || (y-s.Eps < 0)

			if !badX && !badY {
				newParticle := particle.New(
					x, y,

					len(s.particles),

					particle.PARTICLE,
				)

				s.particles = append(s.particles, newParticle)
			}
		}
	}
}
