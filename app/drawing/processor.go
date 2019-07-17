package drawing

import (
	"math"
	"math/cmplx"
	"encoding/json"

	"api/app/drawing/store"
)

type OriginalPoint struct {
	X int
	Y int
	Time float64
}

type Vector struct {
	N int
	Real float64
	Imaginary float64
}

func Process(drawingId int) error {
	//store := store.New()
	originalPoints := createOriginalPoints(drawingId)
	n := 0
	maxVectorCount := 101
	vectors := []Vector{}

	for (len(vectors) < maxVectorCount) {// && vectorsOutsideThreshold(vectors, originalPoints) {
		vectors = append(vectors, buildVector(n, originalPoints))

		if n != 0 {
			vectors = append(vectors, buildVector(n * -1, originalPoints))
		}

		n++
	}

	//store.AddVectors(drawingId, vectors)

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

func vectorsOutsideThreshold(originalPoints []OriginalPoint, vectors []Vector) bool {
	distance := 0.00

	for i := 0; i < len(originalPoints); i++ {
			originalPoint := originalPoints[i]
			time := originalPoints[i].Time
			calculatedVector := calculateVectorSum(time, vectors)
			distance += math.Sqrt(math.Pow(2, calculatedVector.Real - float64(originalPoint.X)) + math.Pow(2, calculatedVector.Imaginary - float64(originalPoint.Y)))
	}

	return distance / float64(len(originalPoints)) < 5
}

func calculateVectorSum(time float64, vectors []Vector) Vector {
	vector := Vector{};

	for i := 0; i < len(vectors); i++ {
		vector.Real += vectors[i].Real;
		vector.Imaginary += vectors[i].Imaginary;
	}

	return vector;
}

func buildVector(n int, originalPoints []OriginalPoint) Vector {
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

	return Vector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func findOriginalPoint(time float64, originalPoints []OriginalPoint) (OriginalPoint, int) {
	for i := 0; i < len(originalPoints); i++ {
		if originalPoints[i].Time >= time {
			return originalPoints[i], i
		}
	}

	return originalPoints[len(originalPoints) - 1], len(originalPoints) - 1
}