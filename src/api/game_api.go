package api

import (
	"net/http"
	"golife/internal"
	"encoding/json"
)

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	print(internal.MATRIX_SIZE)
	// matrix := internal.createFieldMatrix()
	var matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8
	response := struct{
		matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8
	}{
		matrix: matrix,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}