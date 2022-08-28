package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Neighbours struct for the neighbours of each particle
type Neighbours struct {
	Fluid []*particle.Particle
	Bound []*particle.Particle
}

func (s *Simulation) UpdateNeighbours() {
	neighbours := make(map[int]Neighbours)

	for _, i := range s.particles {
		var fluid []*particle.Particle
		var bound []*particle.Particle

		for _, j := range s.FindGridNeighbours(i) {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			if mag < s.H {
				if j.T == particle.Fluid {
					fluid = append(fluid, j)
				} else {
					bound = append(bound, j)
				}
			}
		}

		neighbours[i.Index] = Neighbours{
			Fluid: fluid,
			Bound: bound,
		}
	}

	s.neighbours = neighbours
}
