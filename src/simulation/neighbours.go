package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) UpdateNeighbours() {
	newNeighbours := make(map[int][]*particle.Particle)

	for _, i := range s.particles {
		var neighbours []*particle.Particle

		for _, j := range s.FindGridNeighbours(i) {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			if mag < s.H {
				neighbours = append(neighbours, j)
			}
		}

		newNeighbours[i.Index] = neighbours
	}

	s.neighbours = newNeighbours
}
