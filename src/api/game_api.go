package api

import (
	"encoding/json"
	"fmt"
	"golife/internal"
	"golife/server"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	MAX_MATRIX_ROWS = 100
	MAX_MATRIX_COLS = 100
)

// extract request / responses to another file in the same package
type ApiErrorRes struct {
	Code int
	Msg  string
}

type RenderMatrixReq struct {
	template string
	offsetX  int
	offsetY  int
	rows     int
	cols     int
}

type RenderMatrixRes struct {
	Matrix [][]uint8
}

func (res *RenderMatrixRes) MarshalJson() ([]byte, error) {
	var matrix string
	if res.Matrix == nil {
		matrix = "null"
	} else {
		matrix = strings.Join(strings.Fields(fmt.Sprintf("%d", res.Matrix)), ",")
	}
	jsonResult := fmt.Sprintf(`{"matrix":%s}`, matrix)
	return []byte(jsonResult), nil
}

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	var matrix [][]uint8
	env := server.GetEnv()

	// refactor me
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
	if req_query.Has("rows") {
		request.rows, _ = strconv.Atoi(req_query.Get("rows"))
		if request.rows > MAX_MATRIX_ROWS {
			request.rows = MAX_MATRIX_ROWS
		}
	}
	if req_query.Has("cols") {
		request.cols, _ = strconv.Atoi(req_query.Get("cols"))
		if request.cols > MAX_MATRIX_COLS {
			request.cols = MAX_MATRIX_COLS
		}
	}

	if request.template != "" {
		template_path := filepath.Join(env.Tempalate_folder, request.template)
		matrix = internal.LoadFieldMatrix(template_path, request.offsetX, request.offsetY, request.rows, request.cols)
	} else {
		matrix = internal.CreateFieldMatrix(request.rows, request.cols)
	}
	response := RenderMatrixRes{
		Matrix: matrix,
	}
	json_response, err := response.MarshalJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error marshal: %+v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)
}

type GetNextStepReq struct {
	Matrix [][]uint8 `json:"matrix"`
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

	matrix := internal.NextGeneration(req_body.Matrix)

	response := RenderMatrixRes{
		Matrix: matrix,
	}
	json_response, err := response.MarshalJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error marshal: %+v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)
}
