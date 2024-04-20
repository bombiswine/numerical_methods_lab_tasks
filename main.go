package main

import (
	"fmt"
	"lab_1/matrix"
)

func main() {
	var initialMatrix matrix.Matrix
	err := initialMatrix.ReadFromFile("matrices\\matrix.txt")
	if err != nil {
		fmt.Println("Error: Can't read the matrix", err)
		return
	}

	inverseMatrix, err := initialMatrix.Invert()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = inverseMatrix.WriteToFile("matrices\\inverse_matrix.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Inverse matrix successfully written to inverse_matrix.txt")
}
