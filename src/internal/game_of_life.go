package internal

import (
	"golife/reader"
	"log"
	"strconv"
	"strings"
)

const (
	DEAD_CELL = 0
	LIFE_CELL = 1
)

type Cell struct {
	matrix [][]uint8
	x      int
	y      int
}

func (c *Cell) isAlive() bool {
	return c.matrix[c.x][c.y] == LIFE_CELL
}

func CreateFieldMatrix(rows, cols int) [][]uint8 {
	matrix := make([][]uint8, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]uint8, cols)
	}
	return matrix
}

func LoadFieldMatrix(path string, offset_x, offset_y, rows, cols int) [][]uint8 {
	coordinates_file, err := reader.ReadLines(path)
	if err != nil {
		// fix this. will crash the server if fails
		log.Fatal("error reading file: " + path)
	}
	if offset_x >= rows || offset_y >= cols {
		log.Fatal("offset exceeding matrix. offsetX", offset_x, "offsetY", offset_y)
	}
	matrix := CreateFieldMatrix(rows, cols)
	for _, coord := range coordinates_file {
		life_coord := strings.Split(coord, ",")
		row, _ := strconv.Atoi(life_coord[0])
		col, _ := strconv.Atoi(life_coord[1])
		row += offset_x
		col += offset_y
		if row < rows && col < cols {
			matrix[row][col] = LIFE_CELL
		}
	}
	return matrix
}

func setLifeInCell(c Cell) {
	c.matrix[c.x][c.y] = LIFE_CELL
}

func countCellNeighbours(c Cell) int {
	var i, j, count int

	rows, cols := getMatrixDimensions(c.matrix)
	for i = max(0, c.x-1); i <= min(rows-1, c.x+1); i++ {
		for j = max(0, c.y-1); j <= min(cols-1, c.y+1); j++ {
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

func NextGeneration(matrix [][]uint8) [][]uint8 {
	var n_neighbours int

	rows, cols := getMatrixDimensions(matrix)
	new_matrix := CreateFieldMatrix(rows, cols)
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

func getMatrixDimensions(matrix [][]uint8) (int, int) {
	var rows, cols int

	rows = len(matrix)
	if rows > 0 {
		cols = len(matrix[0])
	}

	return rows, cols
}
