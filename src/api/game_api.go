package api

import (
	"net/http"
	"golife/internal"
	"encoding/json"
)

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	matrix := internal.CreateFieldMatrix()
	response := struct{
		matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8
	}{
		matrix: *matrix,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}