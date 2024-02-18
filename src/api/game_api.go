package api

import (
	"encoding/json"
	"golife/internal"
	"net/http"
)

type RenderResponse struct {
	Matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8
}

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	matrix := internal.CreateFieldMatrix()
	response := RenderResponse{
		Matrix: *matrix,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
