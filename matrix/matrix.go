package matrix

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}

	return &Matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func (matrix *Matrix) ReadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var rows [][]float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		var floatRow []float64
		for _, val := range row {
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			floatRow = append(floatRow, num)
		}
		rows = append(rows, floatRow)
	}

	matrix.Rows = len(rows)
	matrix.Cols = len(rows[0])
	matrix.Data = rows

	return nil
}

func (m *Matrix) WriteToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, row := range m.Data {
		for _, val := range row {
			_, err := writer.WriteString(fmt.Sprintf("%.4f ", val))
			if err != nil {
				return err
			}
		}
		_, err := writer.WriteString("\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}


func (matrix *Matrix) Invert() (*Matrix, error) {
	if !matrix.IsSquare() {
		return nil, errors.New("not a square matrix")
	}

	size := matrix.Rows
	identity := NewMatrix(size, size)
	for i := 0; i < size; i++ {
		identity.Data[i][i] = 1
	}

	// Augment the input matrix with the identity matrix
	augmentedMatrix := NewMatrix(size, size*2)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			augmentedMatrix.Data[i][j] = matrix.Data[i][j]
			augmentedMatrix.Data[i][j+size] = identity.Data[i][j]
		}
	}

	// Perform Gaussian elimination
	for i := 0; i < size; i++ {
		// Find pivot row
		maxIndex := i
		for j := i + 1; j < size; j++ {
			if math.Abs(augmentedMatrix.Data[j][i]) > math.Abs(augmentedMatrix.Data[maxIndex][i]) {
				maxIndex = j
			}
		}
		if maxIndex != i {
			augmentedMatrix.Data[i], augmentedMatrix.Data[maxIndex] = augmentedMatrix.Data[maxIndex], augmentedMatrix.Data[i]
		}

		// normalize the diagonal elements
		divisor := augmentedMatrix.Data[i][i]
		if divisor == 0 {
			return nil, errors.New("singular matrix, inverse does not exist")
		}
		for j := i; j < 2*size; j++ {
			augmentedMatrix.Data[i][j] /= divisor
		}

		// Make other elements in the column 0
		for j := 0; j < size; j++ {
			if j != i {
				ratio := augmentedMatrix.Data[j][i]
				for k := i; k < 2*size; k++ {
					augmentedMatrix.Data[j][k] -= ratio * augmentedMatrix.Data[i][k]
				}
			}
		}
	}

	// Extract the inverse matrix from the augmented matrix
	inverseMatrix := NewMatrix(size, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			inverseMatrix.Data[i][j] = augmentedMatrix.Data[i][j+size]
		}
	}

	return inverseMatrix, nil
}

func (m *Matrix) IsSquare() bool {
	return m.Rows == m.Cols
}

func (thisMatrix *Matrix) Equal(otherMatrix *Matrix) bool {
	const precision float64 = 0.0001

	if thisMatrix.Rows != otherMatrix.Rows || thisMatrix.Cols != otherMatrix.Cols {
		return false
	}

	for i := 0; i < thisMatrix.Rows; i++ {
		for j := 0; j < thisMatrix.Cols; j++ {
			if math.Abs(thisMatrix.Data[i][j]-otherMatrix.Data[i][j]) >= precision {
				return false
			}
		}
	}

	return true
}

func (matrix *Matrix) Product(other *Matrix) *Matrix {
	if matrix.Cols != other.Rows {
		return nil
	}

	result := NewMatrix(matrix.Rows, other.Cols)

	for i := 0; i < matrix.Rows; i++ {
		for j := 0; j < other.Cols; j++ {
			sum := 0.0
			for k := 0; k < matrix.Cols; k++ {
				sum += matrix.Data[i][k] * other.Data[k][j]
			}
			result.Data[i][j] = sum
		}
	}

	return result
}
