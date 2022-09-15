package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/linked_list"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func (s *Simulation) getGridCoords(pos rl.Vector2) (int, int) {
	x := int(math.Floor(float64(pos.X / s.h)))
	y := int(math.Floor(float64(pos.Y / s.h)))

	return x, y
}

func (s *Simulation) initGrid() {
	var newGrid [][]*linked_list.List

	// Initialize grid
	for i := 0; float32(i) < s.rows; i++ {
		var push []*linked_list.List

		for j := 0; float32(j) < s.cols; j++ {
			push = append(push, linked_list.NewList())
		}

		newGrid = append(newGrid, push)
	}

	var remove []int

	// Update grid
	for i, p := range s.particles {
		x, y := s.getGridCoords(p.X)

		xf32 := float32(x)
		yf32 := float32(y)

		if xf32 < 0 || xf32 >= s.cols || yf32 < 0 || yf32 >= s.rows {
			// Flag the particle for deletion if it is outside the grid
			remove = append(remove, i)
		} else {
			// Add the particle to the grid
			node := linked_list.NewNode(p)
			s.nodes[p.Index] = node
			newGrid[y][x].Add(node)
		}
	}

	// Remove particles outside the grid
	for i := len(remove) - 1; i >= 0; i-- {
		r := remove[i]
		s.particles[r] = s.particles[len(s.particles)-1]
		s.particles = s.particles[:len(s.particles)-1]
	}

	s.grid = newGrid
}

func (s *Simulation) appendIfValid(f []*particle.Particle, i *particle.Particle, l *linked_list.List) []*particle.Particle {
	// Loop through the linked list and append the particle if it is a valid neighbour
	list := l.Head

	for list != nil {
		j := list.Value

		rij := rl.Vector2Subtract(i.X, j.X)
		magSq := rl.Vector2LenSqr(rij)

		// Append if particle is in radius and not the particle itself
		if magSq < s.h*s.h && i != j {
			f = append(f, j)
		}

		list = list.Next
	}

	return f
}

func (s *Simulation) findNeighbours(i *particle.Particle) []*particle.Particle {
	x, y := s.getGridCoords(i.X)

	xf32 := float32(x)
	yf32 := float32(y)

	if xf32 < 0 || xf32 >= s.cols || yf32 < 0 || yf32 >= s.rows {
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

func (s *Simulation) updateGridParticle(prev rl.Vector2, p *particle.Particle) {
	prevX, prevY := s.getGridCoords(prev)
	x, y := s.getGridCoords(p.X)

	xf32 := float32(x)
	yf32 := float32(y)

	outOfBounds := xf32 < 0 || xf32 >= s.cols || yf32 < 0 || yf32 >= s.rows
	sameCell := prevX == x && prevY == y

	if !sameCell {
		// Remove the particle from the previous cell
		s.grid[prevY][prevX].Delete(s.nodes[p.Index])

		if !outOfBounds {
			// Add the particle to the new cell
			s.grid[y][x].Add(s.nodes[p.Index])
		}
	}
}
