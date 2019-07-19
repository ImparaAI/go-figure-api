package drawing

import (
	"math"
	"math/cmplx"

	"api/app/drawing/store"
	"api/app/drawing/types"
)

func Process(drawingId int) error {
	//store := store.New()
	originalPoints := createOriginalPoints(drawingId)
	n := 0
	maxDrawVectorCount := 101
	vectors := []types.DrawVector{}

	for (len(vectors) < maxDrawVectorCount) {// && vectorsOutsideThreshold(vectors, originalPoints) {
		vectors = append(vectors, buildDrawVector(n, originalPoints))

		if n != 0 {
			vectors = append(vectors, buildDrawVector(n * -1, originalPoints))
		}

		n++
	}

	//store.AddDrawVectors(drawingId, vectors)

	return nil
}

func createOriginalPoints(drawingId int) []types.OriginalPoint {
	store := store.New()
	drawing := store.Get(drawingId)

	if drawing.Id == 0 {
		panic("The drawing could not be found in storage.")
	}

	normalizeTime(drawing.OriginalPoints)

	return drawing.OriginalPoints
}

func normalizeTime(originalPoints []types.OriginalPoint) {
	finalPoint := originalPoints[len(originalPoints) - 1]

	for i := 0; i < len(originalPoints); i++ {
		originalPoints[i].Time = originalPoints[i].Time / finalPoint.Time
	}
}

func vectorsOutsideThreshold(originalPoints []types.OriginalPoint, vectors []types.DrawVector) bool {
	averageDistance := getAverageDistance(originalPoints, vectors)

	return averageDistance < 5
}

func getAverageDistance(originalPoints []types.OriginalPoint, vectors []types.DrawVector) float64 {
	distance := 0.00

	for i := 0; i < len(originalPoints); i++ {
			original := originalPoints[i]
			estimate := calculateOutput(originalPoints[i].Time, vectors)
			distance += math.Sqrt(math.Pow(2, real(estimate) - float64(original.X)) + math.Pow(2, imag(estimate) - float64(original.Y)))
	}

	return distance / float64(len(originalPoints))
}

func calculateOutput(time float64, vectors []types.DrawVector) complex128 {
	sum := complex(0, 0);

	for i := 0; i < len(vectors); i++ {
		c := complex(vectors[i].Real, vectors[i].Imaginary)
		power := complex(0.00, float64(vectors[i].N) * 2.00 * math.Pi * time)
		sum += c * cmplx.Exp(power)
	}

	return sum;
}

func buildDrawVector(n int, originalPoints []types.OriginalPoint) types.DrawVector {
	time := 0.00
	timeDelta := 0.01
	originalPointsIndex := 0
	cumulativeValue := 0 + 0i
	originalPoint := types.OriginalPoint{}

	for time <= 1 {
		originalPoint, originalPointsIndex = findOriginalPoint(time, originalPoints[originalPointsIndex:])
		originalComplexValue := complex(float64(originalPoint.X), float64(originalPoint.Y))
		cumulativeValue = originalComplexValue * cmplx.Exp(complex(0.00, float64(-n) * 2.0 * math.Pi * time))

		time += timeDelta
	}

	return types.DrawVector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func findOriginalPoint(time float64, originalPoints []types.OriginalPoint) (types.OriginalPoint, int) {
	for i := 0; i < len(originalPoints); i++ {
		if originalPoints[i].Time >= time {
			return originalPoints[i], i
		}
	}

	return originalPoints[len(originalPoints) - 1], len(originalPoints) - 1
}