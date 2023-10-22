package internal

const (
	MATRIX_SIZE = 100
	LIFE_CELL = 1
)

func createFieldMatrix() *[MATRIX_SIZE][MATRIX_SIZE]uint8 {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]uint8 
	return &matrix
}

func setLifeInCell(m *[MATRIX_SIZE][MATRIX_SIZE]uint8, x int, y int) {
	m[x][y] = LIFE_CELL
}