package internal

const (
	MATRIX_SIZE = 100
	DEAD_CELL   = 0
	LIFE_CELL   = 1
)

type Cell struct {
	matrix *[MATRIX_SIZE][MATRIX_SIZE]uint8
	x      int
	y      int
}

func (c *Cell) isAlive() bool {
	return c.matrix[c.x][c.y] == LIFE_CELL
}

func CreateFieldMatrix() *[MATRIX_SIZE][MATRIX_SIZE]uint8 {
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

func NextGeneration(matrix *[MATRIX_SIZE][MATRIX_SIZE]uint8) *[MATRIX_SIZE][MATRIX_SIZE]uint8 {
	var n_neighbours int

	new_matrix := CreateFieldMatrix()
	for i, row := range matrix {
		for j := range row {
			cell := Cell{matrix, i, j}
			n_neighbours = countCellNeighbours(cell)

			if n_neighbours == 3 {
				new_matrix[i][j] = LIFE_CELL
			} else if cell.isAlive() && (n_neighbours == 2 || n_neighbours == 3) {
				new_matrix[i][j] = LIFE_CELL
			} else {
				new_matrix[i][j] = DEAD_CELL
			}
		}
	}

	return new_matrix
}
