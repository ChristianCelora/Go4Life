package api

import (
	"encoding/json"
	"golife/internal"
	"golife/server"
	"net/http"
	"path/filepath"
	"strconv"
)

type ApiErrorRes struct {
	Code int
	Msg  string
}

type RenderMatrixReq struct {
	template string
	offsetX  int
	offsetY  int
}

type RenderMatrixRes struct {
	Matrix [internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8 `json:"matrix"`
}

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	var matrix *[internal.MATRIX_SIZE][internal.MATRIX_SIZE]uint8
	env := server.GetEnv()
	req_query := req.URL.Query()
	request := RenderMatrixReq{
		template: req_query.Get("template"),
	}
	if req_query.Has("offsetX") {
		request.offsetX, _ = strconv.Atoi(req_query.Get("offsetX"))
	}
	if req_query.Has("offsetY") {
		request.offsetY, _ = strconv.Atoi(req_query.Get("offsetY"))
	}

	if request.template != "" {
		template_path := filepath.Join(env.Tempalate_folder, request.template)
		matrix = internal.LoadFieldMatrix(template_path, request.offsetX, request.offsetY)
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
