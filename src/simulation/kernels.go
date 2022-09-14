package simulation

import "math"

func (s *Simulation) poly6(rSq float32) float32 {
	if 0 <= rSq && rSq <= s.h*s.h {
		return s.poly6F * float32(math.Pow(float64(s.h*s.h-rSq), 3))
	} else {
		return 0
	}
}

func (s *Simulation) spikyGrad(r float32) float32 {
	if 0 <= r && r <= s.h {
		return s.spikyGradF * float32(math.Pow(float64(s.h-r), 2))
	} else {
		return 0
	}
}

func (s *Simulation) viscLap(r float32) float32 {
	return s.viscLapF * (s.h - r)
}
