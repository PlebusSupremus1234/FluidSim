package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// VolMapEntry entry for faster volume map values queries
type VolMapEntry struct {
	Vol     float32
	Closest rl.Vector2
}

func (s *Simulation) computeBoundaryVolume(i *particle.Particle) *VolMapEntry {
	var volume float32 = 0

	for _, j := range s.neighbours[i.Index] {
		_, signedDist := boundary.SignedDistance(j.X, s.boundaries)

		extended := boundary.ExtensionFunc(signedDist, s.H, s.CubicSplineF)

		volume += extended
	}

	closest, signedDist := boundary.SignedDistance(i.X, s.boundaries)

	self := boundary.ExtensionFunc(signedDist, s.H, s.CubicSplineF)

	volume += self

	return &VolMapEntry{
		Vol:     volume,
		Closest: closest,
	}
}
