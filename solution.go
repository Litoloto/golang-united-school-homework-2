package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
import (
	"math"
)

type mayType int

const SidesTriangle mayType = 3
const SidesSquare mayType = 4
const SidesCircle mayType = 0

func CalcSquare(sideLen float64, sidesNum mayType) float64 {
	var result float64
	switch sidesNum {
	case SidesCircle:
		result = math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		result = ((math.Sqrt(3)) / 4) * math.Pow(sideLen, 2)
	case SidesSquare:
		result = math.Pow(sideLen, 2)
	default:
		result = 0.0
	}
	return result
}
