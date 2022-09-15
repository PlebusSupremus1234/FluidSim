package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/linked_list"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

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

				node := linked_list.NewNode(newParticle)
				s.nodes[newParticle.Index] = node

				gridX, gridY := s.getGridCoords(newParticle.X)
				s.grid[gridY][gridX].Add(node)
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
