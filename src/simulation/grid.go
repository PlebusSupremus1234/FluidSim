package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/list"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func (s *Simulation) getGridCoords(pos rl.Vector2) (int, int) {
	x := int(math.Floor(float64(pos.X / s.h)))
	y := int(math.Floor(float64(pos.Y / s.h)))

	return x, y
}

func (s *Simulation) outOfBounds(x, y int) bool {
	xf32 := float32(x)
	yf32 := float32(y)

	return xf32 < 0 || xf32 >= s.cols || yf32 < 0 || yf32 >= s.rows
}

func (s *Simulation) initGrid() {
	var newGrid [][]*list.List

	// Initialize grid
	for i := 0; float32(i) < s.rows; i++ {
		var push []*list.List

		for j := 0; float32(j) < s.cols; j++ {
			push = append(push, list.NewList())
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
			// Add particle's node to nodes
			node := list.NewNode(p)
			s.nodes[p.Index] = node

			// Add the particle to the grid
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
