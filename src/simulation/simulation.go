package simulation

import (
	"math"

	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	particles  []*particle.Particle         // Simulation particles
	neighbours map[int][]*particle.Particle // Neighbours for each particle
	grid       [][][]*particle.Particle     // Grid for faster neighbour lookup

	H float32 // Radius

	RestDens  float32 // Rest density
	Stiffness float32 // Stiffness

	Visc float32 // Viscosity coefficient

	Cohe float32 // Cohesion coefficient

	Gravity rl.Vector2 // Gravity
	DT      float32    // Integration timestep

	// Kernel factors
	Poly6F     float32
	SpikyGradF float32
	ViscLapF   float32

	Eps          float32 // Boundary epsilon
	BoundDamping float32 // Boundary damping

	ViewWidth  float32 // View width
	ViewHeight float32 // View height

	COLS float32 // Number of grid columns
	ROWS float32 // Number of grid rows
}

func New(H, cols, rows, width, height float32) *Simulation {
	particles := initParticles(H, cols, rows)

	Hf64 := float64(H)

	return &Simulation{
		particles:  particles,
		neighbours: make(map[int][]*particle.Particle),
		grid:       [][][]*particle.Particle{},

		H: H,

		RestDens:  1000,
		Stiffness: 2000,

		Visc: 200,

		Cohe: 0.5,

		Gravity: rl.NewVector2(0, 9.81),
		DT:      0.0007,

		Poly6F:     4 / float32(math.Pi*math.Pow(Hf64, 8)),
		SpikyGradF: -30 / float32(math.Pi*math.Pow(Hf64, 5)),
		ViscLapF:   40 / float32(math.Pi*math.Pow(Hf64, 5)),

		Eps:          H,
		BoundDamping: -0.5,

		ViewWidth:  width,
		ViewHeight: height,

		COLS: cols,
		ROWS: rows,
	}
}
