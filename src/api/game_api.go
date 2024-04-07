package api

import (
	"encoding/json"
	"golife/internal"
	"net/http"
)

type ApiErrorRes struct {
	Code int
	Msg  string
}

type RenderMatrixRes struct {
	Matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8 `json:"matrix"`
}

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	matrix := internal.CreateFieldMatrix()
	response := RenderMatrixRes{
		Matrix: *matrix,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type GetNextStepReq struct {
	Matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8 `json:"matrix"`
}

func GetNextStep(w http.ResponseWriter, req *http.Request) {
	var req_body GetNextStepReq
	err := json.NewDecoder(req.Body).Decode(&req_body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiErrorRes{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	matrix := internal.NextGeneration(&req_body.Matrix)

	response := RenderMatrixRes{
		Matrix: *matrix,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
