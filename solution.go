package square

import (
	"math"
)

type CustomInt int

const Pi = 3.14

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	var res float64
	switch {
	case sidesNum == 0:
		res = sideLen * sideLen * Pi
	case sidesNum == 4:
		res = sideLen * sideLen
	case sidesNum == 3:
		res = sideLen * sideLen * math.Sqrt(3) / 4
	default:
		return 0
	}
	return res
}
