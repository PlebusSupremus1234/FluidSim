package simulation

import (
	"math"

	"github.com/PlebusSupremus1234/Fluid-Simulation/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) UpdateGrid() {
	var newGrid [][][]*particle.Particle

	for i := 0; float32(i) < s.ROWS; i++ {
		var push [][]*particle.Particle

		for j := 0; float32(j) < s.COLS; j++ {
			push = append(push, []*particle.Particle{})
		}

		newGrid = append(newGrid, push)
	}

	var remove []int

	for i, p := range s.particles {
		x := int(math.Floor(float64(p.X.X / s.H)))
		y := int(math.Floor(float64(p.X.Y / s.H)))

		xf32 := float32(x)
		yf32 := float32(y)

		if xf32 < 0 || xf32 >= s.COLS || yf32 < 0 || yf32 >= s.ROWS {
			// Delete particle if it is outside the grid
			remove = append(remove, i)
		} else {
			newGrid[y][x] = append(newGrid[y][x], p)
		}
	}

	for i := len(remove) - 1; i >= 0; i-- {
		r := remove[i]
		s.particles[r] = s.particles[len(s.particles)-1]
		s.particles = s.particles[:len(s.particles)-1]
	}

	s.grid = newGrid
}

func (s *Simulation) FindGridNeighbours(P *particle.Particle) []*particle.Particle {
	indexX := int(math.Floor(float64(P.X.X / s.H)))
	indexY := int(math.Floor(float64(P.X.Y / s.H)))

	lessX := indexX-1 >= 0
	lessY := indexY-1 >= 0
	moreX := float32(indexX)+1 < s.COLS
	moreY := float32(indexY)+1 < s.ROWS

	var found []*particle.Particle

	if lessX && lessY {
		found = append(found, s.grid[indexY-1][indexX-1]...)
	}
	if lessX {
		found = append(found, s.grid[indexY][indexX-1]...)
	}
	if lessX && moreY {
		found = append(found, s.grid[indexY+1][indexX-1]...)
	}

	if lessY {
		found = append(found, s.grid[indexY-1][indexX]...)
	}
	if moreY {
		found = append(found, s.grid[indexY+1][indexX]...)
	}

	if moreX && lessY {
		found = append(found, s.grid[indexY-1][indexX+1]...)
	}
	if moreX {
		found = append(found, s.grid[indexY][indexX+1]...)
	}
	if moreX && moreY {
		found = append(found, s.grid[indexY+1][indexX+1]...)
	}

	for _, Pi := range s.grid[indexY][indexX] {
		if Pi != P {
			found = append(found, Pi)
		}
	}

	return found
}

func (s *Simulation) UpdateNeighbours() {
	var newNeighbours [][]*particle.Particle

	for range s.particles {
		newNeighbours = append(newNeighbours, []*particle.Particle{})
	}

	for i, Pi := range s.particles {
		neighbours := s.FindGridNeighbours(Pi)

		for _, Pj := range neighbours {
			rij := rl.Vector2Subtract(Pi.X, Pj.X)
			mag := rl.Vector2Length(rij)

			if mag < s.H {
				newNeighbours[i] = append(newNeighbours[i], Pj)
			}
		}
	}

	s.neighbours = newNeighbours
}
