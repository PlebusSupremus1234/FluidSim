package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/list"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) findNeighbours(i *particle.Particle) []*particle.Particle {
	x, y := s.getGridCoords(i.X)

	if s.outOfBounds(x, y) {
		return []*particle.Particle{}
	}

	lessX := x-1 >= 0
	lessY := y-1 >= 0
	moreX := float32(x)+1 < s.cols
	moreY := float32(y)+1 < s.rows

	var found []*particle.Particle

	// Left 3
	if lessX && lessY {
		found = s.appendIfValid(found, i, s.grid[y-1][x-1])
	}
	if lessX {
		found = s.appendIfValid(found, i, s.grid[y][x-1])
	}
	if lessX && moreY {
		found = s.appendIfValid(found, i, s.grid[y+1][x-1])
	}

	// Middle 3
	if lessY {
		found = s.appendIfValid(found, i, s.grid[y-1][x])
	}
	found = s.appendIfValid(found, i, s.grid[y][x])
	if moreY {
		found = s.appendIfValid(found, i, s.grid[y+1][x])
	}

	// Right 3
	if moreX && lessY {
		found = s.appendIfValid(found, i, s.grid[y-1][x+1])
	}
	if moreX {
		found = s.appendIfValid(found, i, s.grid[y][x+1])
	}
	if moreX && moreY {
		found = s.appendIfValid(found, i, s.grid[y+1][x+1])
	}

	return found
}

func (s *Simulation) appendIfValid(f []*particle.Particle, i *particle.Particle, l *list.List) []*particle.Particle {
	// Loop through the linked list and append the particle if it is a valid neighbour

	node := l.Head

	for node != nil {
		j := node.Value

		rij := rl.Vector2Subtract(i.X, j.X)
		magSq := rl.Vector2LenSqr(rij)

		// Append if particle is in radius and not the particle itself
		if magSq < s.h*s.h && i != j {
			f = append(f, j)
		}

		node = node.Next
	}

	return f
}
