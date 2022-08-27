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

	for i, Pi := range s.particles {
		var fluid []*particle.Particle
		var bound []*particle.Particle

		for _, Pj := range s.FindGridNeighbours(Pi) {
			rij := rl.Vector2Subtract(Pi.X, Pj.X)
			mag := rl.Vector2Length(rij)

			if mag <= s.H {
				if Pj.T == particle.Fluid {
					fluid = append(fluid, Pj)
				} else {
					bound = append(bound, Pj)
				}
			}
		}

		neighbours[i] = Neighbours{
			Fluid: fluid,
			Bound: bound,
		}
	}

	s.neighbours = neighbours
}
