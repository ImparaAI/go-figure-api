package drawing

import (
	"math"
	"math/cmplx"
	"encoding/json"

	"api/app/drawing/store"
	"api/app/drawing/types"
)

type OriginalPoint types.OriginalPoint
type DrawVector types.DrawVector

func Process(drawingId int) error {
	//store := store.New()
	originalPoints := createOriginalPoints(drawingId)
	n := 0
	maxDrawVectorCount := 101
	vectors := []DrawVector{}

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

func createOriginalPoints(drawingId int) []OriginalPoint {
	store := store.New()
	drawing, err := store.Get(drawingId)

	if err != nil {
		panic("The drawing could not be found in storage.")
	}

	originalPoints := []OriginalPoint{}
	b := []byte(drawing.OriginalPoints)
	err = json.Unmarshal(b, &originalPoints)

	if err != nil {
		panic("The input points seem to be improperly formatted.")
	}

	normalizeTime(originalPoints)

	return originalPoints
}

func normalizeTime(originalPoints []OriginalPoint) {
	finalPoint := originalPoints[len(originalPoints) - 1]

	for i := 0; i < len(originalPoints); i++ {
		originalPoints[i].Time = originalPoints[i].Time / finalPoint.Time
	}
}

func vectorsOutsideThreshold(originalPoints []OriginalPoint, vectors []DrawVector) bool {
	distance := 0.00

	for i := 0; i < len(originalPoints); i++ {
			originalPoint := originalPoints[i]
			time := originalPoints[i].Time
			estimate := calculateOutput(time, vectors)
			distance += math.Sqrt(math.Pow(2, real(estimate) - float64(originalPoint.X)) + math.Pow(2, imag(estimate) - float64(originalPoint.Y)))
	}

	return distance / float64(len(originalPoints)) < 5
}

func calculateOutput(time float64, vectors []DrawVector) complex128 {
	sum := complex(0, 0);

	for i := 0; i < len(vectors); i++ {
		c := complex(vectors[i].Real, vectors[i].Imaginary)
		power := complex(0.00, float64(vectors[i].N) * 2.00 * math.Pi * time)
		sum += c * cmplx.Exp(power)
	}

	return sum;
}

func buildDrawVector(n int, originalPoints []OriginalPoint) DrawVector {
	time := 0.00
	timeDelta := 0.01
	originalPointsIndex := 0
	cumulativeValue := 0 + 0i
	originalPoint := OriginalPoint{}

	for time <= 1 {
		originalPoint, originalPointsIndex = findOriginalPoint(time, originalPoints[originalPointsIndex:])
		originalComplexValue := complex(float64(originalPoint.X), float64(originalPoint.Y))
		cumulativeValue = originalComplexValue * cmplx.Exp(complex(0.00, float64(-n) * 2.0 * math.Pi * time))

		time += timeDelta
	}

	return DrawVector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func findOriginalPoint(time float64, originalPoints []OriginalPoint) (OriginalPoint, int) {
	for i := 0; i < len(originalPoints); i++ {
		if originalPoints[i].Time >= time {
			return originalPoints[i], i
		}
	}

	return originalPoints[len(originalPoints) - 1], len(originalPoints) - 1
}