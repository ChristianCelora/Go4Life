package internal

import "testing"

func TestCreateFieldMatrix(t *testing.T) {
	n := MATRIX_SIZE
	matrix := CreateFieldMatrix()

	if len(matrix) != n {
		t.Fatalf("Matrix columns count expected %d, actual %d", n, len(matrix))
	}

	if len(matrix[0]) != n {
		t.Fatalf("Matrix rows count expected %d, actual %d", n, len(matrix))
	}
}

func TestSetLifeInCell(t *testing.T) {
	matrix := CreateFieldMatrix()
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
			expected_count: 3,
		},
		{
			other_life_coordinates: []Coordinates{
				Coordinates{x: 0, y: 0},
				Coordinates{x: 0, y: 1},
				Coordinates{x: 0, y: 2},
				Coordinates{x: 1, y: 0},
				Coordinates{x: 1, y: 2},
				Coordinates{x: 2, y: 0},
				Coordinates{x: 2, y: 1},
				Coordinates{x: 2, y: 2},
			}, 
			cell_coordinate: Coordinates{x: 1, y: 1}, 
			expected_count: 8,
		},
		{
			other_life_coordinates: []Coordinates{
				Coordinates{x: 0, y: 0},
				Coordinates{x: 1, y: 2},
				Coordinates{x: 2, y: 0},
				Coordinates{x: 2, y: 1},
			}, 
			cell_coordinate: Coordinates{x: 1, y: 1}, 
			expected_count: 4,
		},
		{
			other_life_coordinates: []Coordinates{
				Coordinates{1,2},
				Coordinates{2,1},
				Coordinates{2,2},
			}, 
			cell_coordinate: Coordinates{x: 1, y: 1}, 
			expected_count: 3,
		},
	}


	for _, test := range tests {
		matrix := CreateFieldMatrix()
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
		matrix_life []Coordinates 
		expected_matrix_life []Coordinates
	} {
		/**
		* xx	xx
		* xx -> xx
		*/
		{	
			matrix_life: []Coordinates{
				Coordinates{1,1},
				Coordinates{1,2},
				Coordinates{2,1},
				Coordinates{2,2},
			},
			expected_matrix_life: []Coordinates{
				Coordinates{1,1},
				Coordinates{1,2},
				Coordinates{2,1},
				Coordinates{2,2},
			},
		},
		/**
		* 		 x
		* xxx -> x
		* 		 x
		*/
		{
			matrix_life: []Coordinates{
				Coordinates{1,0},
				Coordinates{1,1},
				Coordinates{1,2},
			},
			expected_matrix_life: []Coordinates{
				Coordinates{0,1},
				Coordinates{1,1},
				Coordinates{2,1},
			},
		},
		/**
		* 		  x
		* xxx -> xx
		* x		 x
		*/
		{
			matrix_life: []Coordinates{
				Coordinates{1,0},
				Coordinates{1,1},
				Coordinates{1,2},
				Coordinates{2,0},
			},
			expected_matrix_life: []Coordinates{
				Coordinates{0,1},
				Coordinates{1,0},
				Coordinates{1,1},
				Coordinates{2,0},
			},
		},
		/**
		*  x	xx	  
		* xx -> xx
		* x		xx
		*/
		{
			matrix_life: []Coordinates{
				Coordinates{0,1},
				Coordinates{1,0},
				Coordinates{1,1},
				Coordinates{2,0},
			},
			expected_matrix_life: []Coordinates{
				Coordinates{0,0},
				Coordinates{0,1},
				Coordinates{1,0},
				Coordinates{1,1},
				Coordinates{2,0},
				Coordinates{2,1},
			},
		},
		/**
		* xx	 xx 
		* xx -> x  x
		* xx     xx
		*/
		{
			matrix_life: []Coordinates{
				Coordinates{0,1},
				Coordinates{0,2},
				Coordinates{1,1},
				Coordinates{1,2},
				Coordinates{2,1},
				Coordinates{2,2},
			},
			expected_matrix_life: []Coordinates{
				Coordinates{0,1},
				Coordinates{0,2},
				Coordinates{1,0},
				Coordinates{1,3},
				Coordinates{2,1},
				Coordinates{2,2},
			},
		},
	}


	for _, test := range tests {
		old_matrix := CreateFieldMatrix()
		for _, c := range test.matrix_life {
			setLifeInCell(Cell{old_matrix, c.x, c.y})
		}

		expected_matrix := CreateFieldMatrix()
		for _, c := range test.expected_matrix_life {
			setLifeInCell(Cell{expected_matrix, c.x, c.y})
		}

		actual_matrix := nextGeneration(old_matrix)

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