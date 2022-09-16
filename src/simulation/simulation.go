package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/list"
	"math"

	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	particles  []*particle.Particle         // Simulation particles
	neighbours map[int][]*particle.Particle // Particle neighbours
	grid       [][]*list.List               // Grid for faster neighbour lookup
	nodes      map[int]*list.Node           // Linked list nodes for each particle

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

func New(h, cols, rows, width, height float32) *Simulation {
	hf64 := float64(h)

	s := &Simulation{
		particles:  []*particle.Particle{},
		neighbours: make(map[int][]*particle.Particle),
		nodes:      make(map[int]*list.Node),

		h: h,

		rho0:      1000,
		stiffness: -7000,

		nu: 200,

		gravity: rl.NewVector2(0, 9.81),
		dt:      0.0007,

		poly6F:       4 / float32(math.Pi*math.Pow(hf64, 8)),
		spikyGradF:   -30 / float32(math.Pi*math.Pow(hf64, 5)),
		viscLapF:     40 / float32(math.Pi*math.Pow(hf64, 5)),
		cubicSplineF: 40 / (7 * math.Pi * h * h),

		index: 0,

		viewW: width,
		viewH: height,

		cols: cols,
		rows: rows,
	}

	s.initGrid()

	return s
}
