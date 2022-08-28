package simulation

import "math"

func (s *Simulation) Poly6(rSq float32) float32 {
	if 0 <= rSq && rSq <= s.H*s.H {
		return s.Poly6F * float32(math.Pow(float64(s.H*s.H-rSq), 3))
	} else {
		return 0
	}
}

func (s *Simulation) SpikyGrad(r float32) float32 {
	if 0 <= r && r <= s.H {
		return s.SpikyGradF * float32(math.Pow(float64(s.H-r), 2))
	} else {
		return 0
	}
}

func (s *Simulation) ViscLap(r float32) float32 {
	return s.ViscLapF * (s.H - r)
}
