package internal

import (
	"testing"
)

const (
	TEST_MATRIX_SIZE = 20
)

func TestCreateFieldMatrix(t *testing.T) {
	exp_rows := 9
	exp_cols := 4

	matrix := CreateFieldMatrix(exp_rows, exp_cols)
	act_rows, act_cols := getMatrixDimensions(matrix)

	if act_rows != exp_rows {
		t.Fatalf("Matrix rows count expected %d, actual %d", exp_rows, act_rows)
	}

	if act_cols != exp_cols {
		t.Fatalf("Matrix columns count expected %d, actual %d", exp_cols, act_cols)
	}
}

func TestSetLifeInCell(t *testing.T) {
	matrix := CreateFieldMatrix(TEST_MATRIX_SIZE, TEST_MATRIX_SIZE)
	cell := Cell{
		matrix: matrix,
		x:      0,
		y:      0,
	}

	setLifeInCell(cell)

	if matrix[0][0] != LIFE_CELL {
		t.Fatalf("Matrix cell value expected %d, actual %d", LIFE_CELL, matrix[0][0])
	}
}

type Coordinates struct {
	x int
	y int
}

func TestCountCellNeighbours(t *testing.T) {
	var neighbours int

	tests := []struct {
		other_life_coordinates []Coordinates
		cell_coordinate        Coordinates
		expected_count         int
	}{
		{
			other_life_coordinates: []Coordinates{
				{x: 1, y: 0},
				{x: 0, y: 1},
				{x: 1, y: 1},
			},
			cell_coordinate: Coordinates{x: 0, y: 0},
			expected_count:  3,
		},
		{
			other_life_coordinates: []Coordinates{
				{x: 0, y: 0},
				{x: 0, y: 1},
				{x: 0, y: 2},
				{x: 1, y: 0},
				{x: 1, y: 2},
				{x: 2, y: 0},
				{x: 2, y: 1},
				{x: 2, y: 2},
			},
			cell_coordinate: Coordinates{x: 1, y: 1},
			expected_count:  8,
		},
		{
			other_life_coordinates: []Coordinates{
				{x: 0, y: 0},
				{x: 1, y: 2},
				{x: 2, y: 0},
				{x: 2, y: 1},
			},
			cell_coordinate: Coordinates{x: 1, y: 1},
			expected_count:  4,
		},
		{
			other_life_coordinates: []Coordinates{
				{1, 2},
				{2, 1},
				{2, 2},
			},
			cell_coordinate: Coordinates{x: 1, y: 1},
			expected_count:  3,
		},
	}

	for _, test := range tests {
		matrix := CreateFieldMatrix(TEST_MATRIX_SIZE, TEST_MATRIX_SIZE)
		cell := Cell{matrix, test.cell_coordinate.x, test.cell_coordinate.y}
		for _, coordinates := range test.other_life_coordinates {
			setLifeInCell(Cell{matrix: matrix, x: coordinates.x, y: coordinates.y})
		}

		neighbours = countCellNeighbours(cell)

		if neighbours != test.expected_count {
			t.Fatalf("expected %d neighbours (actual %d) for cell %+v", test.expected_count, neighbours, cell)
		}
	}
}

func TestNextGeneration(t *testing.T) {
	tests := []struct {
		matrix_life          []Coordinates
		expected_matrix_life []Coordinates
	}{
		/**
		* xx	xx
		* xx -> xx
		 */
		{
			matrix_life: []Coordinates{
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
			},
			expected_matrix_life: []Coordinates{
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
			},
		},
		/**
		* 		 x
		* xxx -> x
		* 		 x
		 */
		{
			matrix_life: []Coordinates{
				{1, 0},
				{1, 1},
				{1, 2},
			},
			expected_matrix_life: []Coordinates{
				{0, 1},
				{1, 1},
				{2, 1},
			},
		},
		/**
		* 		  x
		* xxx -> xx
		* x		 x
		 */
		{
			matrix_life: []Coordinates{
				{1, 0},
				{1, 1},
				{1, 2},
				{2, 0},
			},
			expected_matrix_life: []Coordinates{
				{0, 1},
				{1, 0},
				{1, 1},
				{2, 0},
			},
		},
		/**
		*  x	xx
		* xx -> xx
		* x		xx
		 */
		{
			matrix_life: []Coordinates{
				{0, 1},
				{1, 0},
				{1, 1},
				{2, 0},
			},
			expected_matrix_life: []Coordinates{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
				{2, 0},
				{2, 1},
			},
		},
		/**
		* xx	 xx
		* xx -> x  x
		* xx     xx
		 */
		{
			matrix_life: []Coordinates{
				{0, 1},
				{0, 2},
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
			},
			expected_matrix_life: []Coordinates{
				{0, 1},
				{0, 2},
				{1, 0},
				{1, 3},
				{2, 1},
				{2, 2},
			},
		},
	}

	for _, test := range tests {
		old_matrix := CreateFieldMatrix(TEST_MATRIX_SIZE, TEST_MATRIX_SIZE)
		for _, c := range test.matrix_life {
			setLifeInCell(Cell{old_matrix, c.x, c.y})
		}

		expected_matrix := CreateFieldMatrix(TEST_MATRIX_SIZE, TEST_MATRIX_SIZE)
		for _, c := range test.expected_matrix_life {
			setLifeInCell(Cell{expected_matrix, c.x, c.y})
		}

		actual_matrix := NextGeneration(old_matrix)

		for i, row := range actual_matrix {
			for j, cell := range row {
				if cell != expected_matrix[i][j] {
					t.Logf("matrix: %v", actual_matrix)
					t.Fatalf("cell %v value %d, different than expected %d", Cell{actual_matrix, i, j}, cell, expected_matrix[i][j])
				}
			}
		}
	}
}

func TestLoadFieldMatrix(t *testing.T) {
	tests := []struct {
		pattern_path         string
		offset_x             int
		offset_y             int
		expected_matrix_life []Coordinates
	}{
		{
			pattern_path: "../templates/test-square",
			offset_x:     0,
			offset_y:     0,
			expected_matrix_life: []Coordinates{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
			},
		},
		{
			pattern_path: "../templates/test-square",
			offset_x:     5,
			offset_y:     0,
			expected_matrix_life: []Coordinates{
				{5, 0},
				{5, 1},
				{6, 0},
				{6, 1},
			},
		},
		{
			pattern_path: "../templates/test-square",
			offset_x:     0,
			offset_y:     4,
			expected_matrix_life: []Coordinates{
				{0, 4},
				{0, 5},
				{1, 4},
				{1, 5},
			},
		},
		{
			pattern_path: "../templates/test-square",
			offset_x:     7,
			offset_y:     10,
			expected_matrix_life: []Coordinates{
				{7, 10},
				{7, 11},
				{8, 10},
				{8, 11},
			},
		},
	}

	for _, test := range tests {
		actual_matrix := LoadFieldMatrix(test.pattern_path, test.offset_x, test.offset_y, TEST_MATRIX_SIZE, TEST_MATRIX_SIZE)
		expected_matrix := CreateFieldMatrix(TEST_MATRIX_SIZE, TEST_MATRIX_SIZE)
		for _, c := range test.expected_matrix_life {
			setLifeInCell(Cell{expected_matrix, c.x, c.y})
		}

		for i, row := range actual_matrix {
			for j, cell := range row {
				if cell != expected_matrix[i][j] {
					t.Logf("matrix: %v", actual_matrix)
					t.Fatalf("cell %v value %d, different than expected %d", Cell{actual_matrix, i, j}, cell, expected_matrix[i][j])
				}
			}
		}
	}
}
