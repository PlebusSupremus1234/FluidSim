package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/list"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
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

	// Update grid
	node := s.particles.Head

	for node != nil {
		// Add particle's node to grid nodes
		gridNode := list.NewNode(node.Value)
		s.gridNodes[node.Value.Index] = gridNode

		// Add the particle to the grid
		x, y := s.getGridCoords(node.Value.X)
		newGrid[y][x].Add(gridNode)

		node = node.Next
	}

	s.grid = newGrid
}

func (s *Simulation) updateGridParticle(prev rl.Vector2, p *particle.Particle) {
	prevX, prevY := s.getGridCoords(prev)
	x, y := s.getGridCoords(p.X)

	sameCell := prevX == x && prevY == y
	if !sameCell {
		// Remove the particle from the previous cell
		s.grid[prevY][prevX].Delete(s.gridNodes[p.Index])

		if !s.outOfBounds(x, y) {
			// Add the particle to the new cell
			s.grid[y][x].Add(s.gridNodes[p.Index])
		} else {
			delete(s.gridNodes, p.Index) // Remove the grid node

			s.particles.Delete(s.particleNodes[p.Index]) // Remove the particle from the simulation

			delete(s.particleNodes, p.Index) // Remove the particle's node
			delete(s.neighbours, p.Index)    // Remove the particle's neighbours
		}
	}
}
