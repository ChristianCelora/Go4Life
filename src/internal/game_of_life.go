package internal

import (
	"golife/reader"
	"log"
	"strconv"
	"strings"
)

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

func LoadFieldMatrix(path string) *[MATRIX_SIZE][MATRIX_SIZE]uint8 {
	coordinates_file, err := reader.ReadLines(path)
	if err != nil {
		// fix this. will crash the server if fails
		log.Fatal("error reading file: " + path)
	}
	var matrix [MATRIX_SIZE][MATRIX_SIZE]uint8
	for _, coord := range coordinates_file {
		life_coord := strings.Split(coord, ",")
		row, _ := strconv.Atoi(life_coord[0])
		col, _ := strconv.Atoi(life_coord[1])
		matrix[row][col] = LIFE_CELL
	}
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
