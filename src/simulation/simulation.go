package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	"math"

	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	particles []*particle.Particle     // Simulation particles
	grid      [][][]*particle.Particle // Grid for faster neighbour lookup

	boundaries []*boundary.Boundary // Boundaries
	volumeMap  map[int]*VolMapEntry // Volume map

	h float32 // Radius

	rho0      float32 // Rest density
	stiffness float32 // Stiffness

	nu float32 // Viscosity coefficient

	gravity rl.Vector2 // Gravity
	dt      float32    // Timestep

	// Kernel factors
	poly6F       float32
	spikyGradF   float32
	viscLapF     float32
	cubicSplineF float32

	index int // Latest particle index

	viewW float32 // View width
	viewH float32 // View height

	// Number of grid columns and rows
	cols float32
	rows float32
}

func New(H, cols, rows, width, height float32) *Simulation {
	Hf64 := float64(H)

	return &Simulation{
		particles: []*particle.Particle{},
		grid:      [][][]*particle.Particle{},

		boundaries: initBoundaries(width, height, H),
		volumeMap:  make(map[int]*VolMapEntry),

		h: H,

		rho0:      1000,
		stiffness: -7000,

		nu: 200,

		gravity: rl.NewVector2(0, 9.81),
		dt:      0.0007,

		poly6F:       4 / float32(math.Pi*math.Pow(Hf64, 8)),
		spikyGradF:   -30 / float32(math.Pi*math.Pow(Hf64, 5)),
		viscLapF:     40 / float32(math.Pi*math.Pow(Hf64, 5)),
		cubicSplineF: 40 / (7 * math.Pi * H * H),

		index: -1,

		viewW: width,
		viewH: height,

		cols: cols,
		rows: rows,
	}
}
