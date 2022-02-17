package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myType int

const (
	SidesCircle   = myType(0)
	SidesTriangle = myType(3)
	SidesSquare   = myType(4)
)

func CalcSquare(sideLen float64, sidesNum myType) float64 {
	ans := 0.
	switch sidesNum {
	case SidesCircle:
		ans = math.Pi * sideLen * sideLen
	case SidesTriangle:
		ans = math.Sqrt(3) * sideLen * sideLen / 4
	case SidesSquare:
		ans = sideLen * sideLen
	}
	return ans
}
