package api

import (
	"encoding/json"
	"golife/internal"
	"golife/server"
	"net/http"
	"path/filepath"
)

type ApiErrorRes struct {
	Code int
	Msg  string
}

type RenderMatrixReq struct {
	Template string
}

type RenderMatrixRes struct {
	Matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8 `json:"matrix"`
}

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	var template string
	var matrix *[internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8
	env := server.GetEnv()
	template = req.URL.Query().Get("Template")

	if template != "" {
		template_path := filepath.Join(env.Tempalate_folder, template)
		matrix = internal.LoadFieldMatrix(template_path)
	} else {
		matrix = internal.CreateFieldMatrix()
	}
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
