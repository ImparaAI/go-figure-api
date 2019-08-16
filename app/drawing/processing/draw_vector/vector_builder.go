package draw_vector

import (
	"math"
	"math/cmplx"

	"api/app/drawing/types"
	"api/app/util"
)

type VectorBuilder struct {
	currentOriginalPointsIndex int
	originalPoints             []types.OriginalPoint
}

func (vectorBuilder *VectorBuilder) Build(n int, providedOriginalPoints []types.OriginalPoint) types.DrawVector {
	vectorBuilder.originalPoints = providedOriginalPoints
	currentTime := 0.00
	timeDelta := 0.001
	vectorBuilder.currentOriginalPointsIndex = 0
	cumulativeValue := 0 + 0i
	originalPoint := types.OriginalPoint{}

	for util.FloatCompare(currentTime, 1.00, 0.0001) < 0 {
		originalPoint = vectorBuilder.findOriginalPoint(currentTime)
		originalComplexValue := complex(float64(originalPoint.X), float64(originalPoint.Y))
		cumulativeValue += originalComplexValue * cmplx.Exp(complex(0.00, float64(-n)*2.0*math.Pi*currentTime)) * complex(timeDelta, 0)

		currentTime += timeDelta
	}

	return types.DrawVector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func (vectorBuilder *VectorBuilder) findOriginalPoint(time float64) types.OriginalPoint {
	var originalPoint types.OriginalPoint

	for i := vectorBuilder.currentOriginalPointsIndex; i < len(vectorBuilder.originalPoints); i++ {
		originalPoint = vectorBuilder.originalPoints[i]

		if util.FloatCompare(originalPoint.Time, time, 0.001) == 0 {
			vectorBuilder.currentOriginalPointsIndex = i

			return originalPoint
		}

		if i != 0 && vectorBuilder.timeBetweenPoints(time, vectorBuilder.originalPoints[i-1], originalPoint) {
			p1 := vectorBuilder.originalPoints[i-1]
			p2 := originalPoint

			vectorBuilder.currentOriginalPointsIndex = i - 1

			return types.OriginalPoint{
				Time: time,
				X:    int(vectorBuilder.getLinearAverage(time, p1.Time, p2.Time, p1.X, p2.X)),
				Y:    int(vectorBuilder.getLinearAverage(time, p1.Time, p2.Time, p1.Y, p2.Y)),
			}
		}
	}

	vectorBuilder.currentOriginalPointsIndex = len(vectorBuilder.originalPoints) - 1

	return vectorBuilder.originalPoints[len(vectorBuilder.originalPoints)-1]
}

func (vectorBuilder *VectorBuilder) timeBetweenPoints(time float64, p1, p2 types.OriginalPoint) bool {
	return util.FloatCompare(time, p1.Time, 0.001) == 1 && util.FloatCompare(time, p2.Time, 0.001) == -1
}

func (vectorBuilder *VectorBuilder) getLinearAverage(input float64, x1, x2 float64, y1, y2 int) float64 {
	slope := float64(y2-y1) / (x2 - x1)
	intercept := float64(y2) - slope*x2

	return slope*input + intercept
}
