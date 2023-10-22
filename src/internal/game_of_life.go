package internal

const (
	MATRIX_SIZE = 100
)

func createFieldMatrix() *[MATRIX_SIZE][MATRIX_SIZE]uint8 {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]uint8 
	return &matrix
}