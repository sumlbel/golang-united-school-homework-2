package square

import "math"

type NumberOfSides uint8

const (
	SidesTriangle NumberOfSides = 3
  SidesSquare NumberOfSides = 4
  SidesCircle NumberOfSides = 0
)

func CalcSquare(sideLen float64, sidesNum NumberOfSides) float64 {
	switch sidesNum {
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	case SidesTriangle:
		return math.Sqrt(3.0) * sideLen * sideLen / 4.0 
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0.0
	}
}
