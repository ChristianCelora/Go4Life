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

func countCellNeighbours(c Cell) int {
	var i, j, count int


	for i = max(0, c.x-1); i <= min(MATRIX_SIZE-1, c.x+1); i++ {
		for j = max(0, c.y-1); j <= min(MATRIX_SIZE-1, c.y+1); j++ {
			if i == c.x && j == c.y {
				continue
			}
			if c.matrix[i][j] == LIFE_CELL {
				count++
			}
		}
	}

	return count
}