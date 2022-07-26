package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/list"
	"math"

	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	particles  *list.List                   // Simulation particles
	neighbours map[int][]*particle.Particle // Particle neighbours
	grid       [][]*list.List               // Grid for faster neighbour lookup

	particleNodes map[int]*list.Node // Linked list particle nodes for each particle
	gridNodes     map[int]*list.Node // Linked list grid nodes for each particle

	h float32 // Radius

	rho0      float32 // Rest density
	stiffness float32 // Stiffness

	nu float32 // Viscosity coefficient

	gravity rl.Vector2 // Gravity
	dt      float32    // Timestep

	// Kernel factors
	poly6F     float32
	spikyGradF float32
	viscLapF   float32

	index int // Latest particle index

	scale float32 // Scale factor for the simulation

	simW float32 // Simulation width
	simH float32 // Simulation height

	viewW float32 // View width
	viewH float32 // View height

	// Number of grid columns and rows
	cols float32
	rows float32
}

func New(h, simW, simH, viewW, viewH float32) *Simulation {
	hf64 := float64(h)

	s := &Simulation{
		particles:  list.NewList(),
		neighbours: make(map[int][]*particle.Particle),

		particleNodes: make(map[int]*list.Node),
		gridNodes:     make(map[int]*list.Node),

		h: h,

		rho0:      0.014920775,
		stiffness: 7000,

		nu: 50,

		gravity: rl.NewVector2(0, 9.81),
		dt:      0.0007,

		poly6F:     4 / float32(math.Pi*math.Pow(hf64, 8)),
		spikyGradF: -30 / float32(math.Pi*math.Pow(hf64, 5)),
		viscLapF:   40 / float32(math.Pi*math.Pow(hf64, 5)),

		index: 0,

		scale: viewW / simW,

		simW: simW,
		simH: simH,

		viewW: viewW,
		viewH: viewH,

		cols: float32(math.Ceil(float64(simW / h))),
		rows: float32(math.Ceil(float64(simH / h))),
	}

	s.initGrid()

	return s
}
