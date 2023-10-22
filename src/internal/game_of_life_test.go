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