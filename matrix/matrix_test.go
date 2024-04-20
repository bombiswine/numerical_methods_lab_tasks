package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInverse(t *testing.T) {
	testCases := []struct {
		name        string
		inputMatrix [][]float64
		expErrMsg   string
	}{
		{
			name: "test 1",
			inputMatrix: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expErrMsg: "",
		}, {
			name: "test 2",
			inputMatrix: [][]float64{
				{2, 0, 0},
				{0, 2, 0},
				{0, 0, 2},
			},
			expErrMsg: "",
		}, {
			name: "test 3",
			inputMatrix: [][]float64{
				{2, 1, 0},
				{0, 2, 0},
				{0, 0, 4},
			},
			expErrMsg: "",
		}, {
			name: "test 4",
			inputMatrix: [][]float64{
				{2, 4, 6},
				{-1, 3.5, 1},
				{0.5, 0, 4},
			},
			expErrMsg: "",
		}, {
			name: "test 5",
			inputMatrix: [][]float64{
				{2, 1, 5},
				{1, 0, 0},
			},
			expErrMsg: "not a square matrix",
		},
	}

	assert := assert.New(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			initialMatrix := NewMatrix(len(tc.inputMatrix), len(tc.inputMatrix[0]))
			initialMatrix.Data = tc.inputMatrix

			actualInverse, err := initialMatrix.Invert()

			// create the identity matrix to test if the inverted matrix found correctly
			size := len(tc.inputMatrix)
			identity := NewMatrix(size, size)
			for i := 0; i < size; i++ {
				identity.Data[i][i] = 1
			}

			// actual test
			if tc.expErrMsg == "" {
				assert.NoError(err, "unexpected error")
				assert.True(identity.Equal(actualInverse.Product(initialMatrix)), "invalid inverse matrix")
			} else {
				assert.Error(err, "expected error")
				assert.EqualError(err, tc.expErrMsg, "unexpected error message")
				assert.Nil(actualInverse, "expected nil result for error case")
			}
		})
	}
}
