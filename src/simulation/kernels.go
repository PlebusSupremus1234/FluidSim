package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func (s *Simulation) poly6(r rl.Vector2) float32 {
	magSq := rl.Vector2LenSqr(r)

	if 0 <= magSq && magSq <= s.h*s.h {
		return s.poly6F * float32(math.Pow(float64(s.h*s.h-magSq), 3))
	} else {
		return 0
	}
}

func (s *Simulation) spikyGrad(r rl.Vector2) rl.Vector2 {
	mag := rl.Vector2Length(r)

	if 0 <= mag && mag <= s.h {
		factor := s.spikyGradF * float32(math.Pow(float64(s.h-mag), 2))
		return rl.Vector2Scale(rl.Vector2Normalize(r), factor)
	} else {
		return rl.Vector2Zero()
	}
}

func (s *Simulation) viscLap(r rl.Vector2) float32 {
	mag := rl.Vector2Length(r)

	if 0 <= mag && mag <= s.h {
		return s.viscLapF * (s.h - mag)
	} else {
		return 0
	}
}
