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
	var boundVol float32 = 0

	for _, j := range s.findNeighbours(i) {
		_, signedDist := boundary.SignedDistance(j.X, s.boundaries)

		extended := boundary.Extend(signedDist, s.h, s.cubicSplineF)
		boundVol += extended
		//boundVol += j.M / j.Rho * extended
	}

	closest, signedDist := boundary.SignedDistance(i.X, s.boundaries)

	self := boundary.Extend(signedDist, s.h, s.cubicSplineF)
	boundVol += self
	//boundVol += i.M / i.Rho * self

	return &VolMapEntry{
		Vol:     boundVol,
		Closest: closest,
	}
}
