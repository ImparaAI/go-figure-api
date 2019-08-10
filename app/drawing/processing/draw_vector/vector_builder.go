package draw_vector

import (
	"math"
	"math/cmplx"

	"api/app/drawing/types"
	"api/app/util"
)

var currentOriginalPointsIndex int = 0
var originalPoints []types.OriginalPoint

func buildDrawVector(n int, providedOriginalPoints []types.OriginalPoint) types.DrawVector {
	originalPoints = providedOriginalPoints
	currentTime := 0.00
	timeDelta := 0.001
	currentOriginalPointsIndex = 0
	cumulativeValue := 0 + 0i
	originalPoint := types.OriginalPoint{}

	for util.FloatCompare(currentTime, 1.00, 0.0001) < 0 {
		originalPoint = findOriginalPoint(currentTime)
		originalComplexValue := complex(float64(originalPoint.X), float64(originalPoint.Y))
		cumulativeValue += originalComplexValue * cmplx.Exp(complex(0.00, float64(-n)*2.0*math.Pi*currentTime)) * complex(timeDelta, 0)

		currentTime += timeDelta
	}

	return types.DrawVector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func findOriginalPoint(time float64) types.OriginalPoint {
	var originalPoint types.OriginalPoint

	for i := currentOriginalPointsIndex; i < len(originalPoints); i++ {
		originalPoint = originalPoints[i]

		if util.FloatCompare(originalPoint.Time, time, 0.001) == 0 {
			currentOriginalPointsIndex = i

			return originalPoint
		}

		if i != 0 && timeBetweenPoints(time, &originalPoints[i-1], &originalPoint) {
			p1 := originalPoints[i-1]
			p2 := originalPoint

			currentOriginalPointsIndex = i - 1

			return types.OriginalPoint{
				Time: time,
				X:    int(getLinearAverage(time, p1.Time, p2.Time, p1.X, p2.X)),
				Y:    int(getLinearAverage(time, p1.Time, p2.Time, p1.Y, p2.Y)),
			}
		}
	}

	currentOriginalPointsIndex = len(originalPoints) - 1

	return originalPoints[len(originalPoints)-1]
}

func timeBetweenPoints(time float64, p1, p2 *types.OriginalPoint) bool {
	return util.FloatCompare(time, p1.Time, 0.001) == 1 && util.FloatCompare(time, p2.Time, 0.001) == -1
}

func getLinearAverage(input float64, x1, x2 float64, y1, y2 int) float64 {
	slope := float64(y2-y1) / (x2 - x1)
	intercept := float64(y2) - slope*x2

	return slope*input + intercept
}
