package internal

const (
	MATRIX_SIZE = 100
	LIFE_CELL = 1
)

type Cell struct {
	matrix *[MATRIX_SIZE][MATRIX_SIZE]uint8
	x int
	y int
}

func createFieldMatrix() *[MATRIX_SIZE][MATRIX_SIZE]uint8 {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]uint8 
	return &matrix
}

func setLifeInCell(c Cell) {
	c.matrix[c.x][c.y] = LIFE_CELL
}