package processing

import (
	"math"
	"math/cmplx"

	"api/app/util"
	"api/app/drawing/types"
)

func buildDrawVectors(originalPoints []types.OriginalPoint) []types.DrawVector {
	n := 0
	vectors := []types.DrawVector{}

	for (len(vectors) < 100) && !vectorsAproximateOriginal(vectors, originalPoints) {
		vectors = append(vectors, buildDrawVector(n, originalPoints))
		n = getNextN(n);
	}

	return vectors
}

func getNextN(n int) int {
	if n > 0 {
		return -1 * n
	}

	return -1 * n + 1
}

func vectorsAproximateOriginal(vectors []types.DrawVector, originalPoints []types.OriginalPoint) bool {
	averageDistance := getAverageDistance(originalPoints, vectors)

	return averageDistance < 5
}

func getAverageDistance(originalPoints []types.OriginalPoint, vectors []types.DrawVector) float64 {
	distance := 0.00

	for i := 0; i < len(originalPoints); i++ {
		original := originalPoints[i]
		estimate := calculateOutput(originalPoints[i].Time, vectors)
		distance += math.Sqrt(math.Pow(real(estimate) - float64(original.X), 2) + math.Pow(imag(estimate) - float64(original.Y), 2))
	}

	return distance / float64(len(originalPoints))
}

func calculateOutput(time float64, vectors []types.DrawVector) complex128 {
	sum := complex(0, 0);

	for _, vector := range vectors {
		c := complex(vector.Real, vector.Imaginary)
		power := complex(0.00, float64(vector.N) * 2.00 * math.Pi * time)
		sum += c * cmplx.Exp(power)
	}

	return sum;
}

func buildDrawVector(n int, originalPoints []types.OriginalPoint) types.DrawVector {
	time := 0.00
	timeDelta := 0.001
	originalPointsIndex := 0
	cumulativeValue := 0 + 0i
	originalPoint := types.OriginalPoint{}

	for util.FloatCompare(time, 1.00, 0.0001) < 0 {
		originalPoint, originalPointsIndex = findOriginalPoint(time, originalPoints[originalPointsIndex:])
		originalComplexValue := complex(float64(originalPoint.X), float64(originalPoint.Y))
		cumulativeValue += originalComplexValue * cmplx.Exp(complex(0.00, float64(-n) * 2.0 * math.Pi * time)) * complex(timeDelta, 0)

		time += timeDelta
	}

	return types.DrawVector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func findOriginalPoint(time float64, originalPoints []types.OriginalPoint) (types.OriginalPoint, int) {
	for i, originalPoint := range originalPoints {
		if util.FloatCompare(originalPoint.Time, time, 0.001) == 0 {
			return originalPoint, i
		}

		if i != 0 && timeBetweenPoints(time, originalPoints[i-1] , originalPoint) {
			p1 := originalPoints[i-1]
			p2 := originalPoint

			return types.OriginalPoint{
				Time: time,
				X: int(getLinearAverage(time, p1.Time, p2.Time, p1.X, p2.X)),
				Y: int(getLinearAverage(time, p1.Time, p2.Time, p1.Y, p2.Y)),
			}, i - 1
		}
	}

	return originalPoints[len(originalPoints) - 1], len(originalPoints) - 1
}

func timeBetweenPoints(time float64, p1, p2 types.OriginalPoint) bool {
	return util.FloatCompare(time, p1.Time, 0.001) == 1 && util.FloatCompare(time, p2.Time, 0.001) == -1
}

func getLinearAverage(input float64, x1, x2 float64, y1, y2 int) float64 {
	slope := float64(y2 - y1) / (x2 - x1)
	intercept := float64(y2) - slope * x2

	return slope * input + intercept
}