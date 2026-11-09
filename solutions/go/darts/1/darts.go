package darts

import "math"

func Score(x, y float64) int {
    radius := math.Sqrt(x * x + y * y)
    if radius <= 10 {
		if radius <= 5 {
			if radius <= 1 {
                return 10
            }
            return 5
        }
        return 1
    }
	return 0
}
