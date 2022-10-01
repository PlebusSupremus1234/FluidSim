package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/list"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

func (s *Simulation) SpawnParticles() {
	// Spawn particles in a 10x10 grid around the mouse

	for i := -5; i < 5; i++ {
		for j := -5; j < 5; j++ {
			mousePos := rl.GetMousePosition()
			mX, mY := mousePos.X/s.scale, mousePos.Y/s.scale

			x := mX + float32(j)*s.h + rand.Float32()/2
			y := mY + float32(i)*s.h + rand.Float32()/2

			badX := (x+s.h > s.simW) || (x-s.h < 0)
			badY := (y+s.h > s.simH) || (y-s.h < 0)

			// Spawn the particle if its in bounds
			if !badX && !badY {
				newParticle := particle.New(x, y, s.index)

				// Create grid node for the particle
				node := list.NewNode(newParticle)
				s.gridNodes[newParticle.Index] = node

				// Create a particle node to the particles list
				particleNode := list.NewNode(newParticle)
				s.particleNodes[newParticle.Index] = particleNode
				s.particles.Add(particleNode)
				s.index++

				// Add node to the grid
				gridX, gridY := s.getGridCoords(newParticle.X)
				s.grid[gridY][gridX].Add(node)
			}
		}
	}
}
