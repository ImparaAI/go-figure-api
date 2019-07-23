package test

import (
	"math"
	"testing"
	"math/cmplx"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"api/app/util"
	"api/app/drawing/types"
	"api/app/drawing/store"
	"api/app/drawing/processing"
)

func TestCircle(t *testing.T) {
	store := store.New()
	radius := 100.00
	id := store.Create(buildUnitCirclePoints(radius))
	processing.Process(id)
	result := store.Get(id)

	expected := []types.DrawVector{
		types.DrawVector{N: 0, Real: 0.00, Imaginary: 0.00},
		types.DrawVector{N: 1, Real: radius, Imaginary: 0.00},
	}

	assert.Equal(t, 2, len(result.DrawVectors))
	assert.Equal(t, true, cmp.Equal(expected, result.DrawVectors, getVectorComparer(0.8)))
}

func buildUnitCirclePoints(radius float64) ([]types.OriginalPoint) {
	points := []types.OriginalPoint{}

	for t := 0.00; util.FloatCompare(t, 1.00, 0.0001) <= 0; t += 0.001 {
		vector := cmplx.Exp(complex(0.00, 2.00 * math.Pi * t))
		points = append(points, types.OriginalPoint{
			Time: t,
			X: int(real(vector) * radius),
			Y: int(imag(vector) * radius),
		})
	}

	return points
}

func getVectorComparer(tolernce float64) cmp.Option {
	return cmp.Comparer(func(a, b types.DrawVector) bool {
		if a.N != b.N {
			return false
		}

		if math.Abs(a.Real - b.Real) > tolernce {
			return false
		}

		if math.Abs(a.Imaginary - b.Imaginary) > tolernce {
			return false
		}

		return true;
	})
}