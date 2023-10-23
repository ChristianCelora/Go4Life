package internal

import "testing"

func TestCreateFieldMatrix(t *testing.T) {
	n := MATRIX_SIZE
	matrix := createFieldMatrix()

	if len(matrix) != n {
		t.Fatalf("Matrix columns count expected %d, actual %d", n, len(matrix))
	}

	if len(matrix[0]) != n {
		t.Fatalf("Matrix rows count expected %d, actual %d", n, len(matrix))
	}
}

func TestSetLifeInCell(t *testing.T) {
	matrix := createFieldMatrix()
	cell := Cell{
		matrix: matrix, 
		x: 0, 
		y: 0,
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
		cell_coordinate Coordinates 
		expected_count int
	} {
		{
			other_life_coordinates: []Coordinates{
				Coordinates{x: 1, y: 0},
				Coordinates{x: 0, y: 1},
				Coordinates{x: 1, y: 1},
			}, 
			cell_coordinate: Coordinates{x: 0, y: 0}, 
			expected_count: 2,
		},
	}


	for _, test := range tests {
		matrix := createFieldMatrix()
		cell := Cell{matrix, test.cell_coordinate.x, test.cell_coordinate.y}
		for _, coordinates := range test.other_life_coordinates {
			setLifeInCell(Cell{matrix: matrix, x: coordinates.x, y: coordinates.y})
		}
		
		neighbours = countCellNeighbours(cell)
		
		if neighbours != 3 {
			t.Fatalf("expected %d neighbours for cell %+v", test.expected_count, cell)
		}
	}
}