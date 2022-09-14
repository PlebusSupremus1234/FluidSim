package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func (s *Simulation) updateGrid() {
	var newGrid [][][]*particle.Particle

	// Initialize grid
	for i := 0; float32(i) < s.rows; i++ {
		var push [][]*particle.Particle

		for j := 0; float32(j) < s.cols; j++ {
			push = append(push, []*particle.Particle{})
		}

		newGrid = append(newGrid, push)
	}

	var remove []int

	// Update grid
	for i, p := range s.particles {
		x := int(math.Floor(float64(p.X.X / s.h)))
		y := int(math.Floor(float64(p.X.Y / s.h)))

		xf32 := float32(x)
		yf32 := float32(y)

		if xf32 < 0 || xf32 >= s.cols || yf32 < 0 || yf32 >= s.rows {
			// Flag the particle for deletion if it is outside the grid
			remove = append(remove, i)
		} else {
			// Add the particle to the grid
			newGrid[y][x] = append(newGrid[y][x], p)
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

func (s *Simulation) findNeighbours(i *particle.Particle) []*particle.Particle {
	indexX := int(math.Floor(float64(i.X.X / s.h)))
	indexY := int(math.Floor(float64(i.X.Y / s.h)))

	lessX := indexX-1 >= 0
	lessY := indexY-1 >= 0
	moreX := float32(indexX)+1 < s.cols
	moreY := float32(indexY)+1 < s.rows

	var found []*particle.Particle

	// Left 3
	if lessX && lessY {
		found = append(found, s.grid[indexY-1][indexX-1]...)
	}
	if lessX {
		found = append(found, s.grid[indexY][indexX-1]...)
	}
	if lessX && moreY {
		found = append(found, s.grid[indexY+1][indexX-1]...)
	}

	// Middle 3
	if lessY {
		found = append(found, s.grid[indexY-1][indexX]...)
	}
	found = append(found, s.grid[indexY][indexX]...)
	if moreY {
		found = append(found, s.grid[indexY+1][indexX]...)
	}

	// Right 3
	if moreX && lessY {
		found = append(found, s.grid[indexY-1][indexX+1]...)
	}
	if moreX {
		found = append(found, s.grid[indexY][indexX+1]...)
	}
	if moreX && moreY {
		found = append(found, s.grid[indexY+1][indexX+1]...)
	}

	var neighbours []*particle.Particle

	// Remove particles outside the radius and the particle itself
	for _, j := range found {
		rij := rl.Vector2Subtract(i.X, j.X)
		magSq := rl.Vector2LenSqr(rij)

		if magSq < s.h*s.h && i != j {
			neighbours = append(neighbours, j)
		}
	}

	return neighbours
}
