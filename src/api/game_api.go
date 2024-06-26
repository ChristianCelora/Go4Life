package api

import (
	"encoding/json"
	"golife/internal"
	"golife/server"
	"log/slog"
	"net/http"
	"path/filepath"
)

func RenderMatrix(w http.ResponseWriter, req *http.Request) {
	var matrix [][]uint8

	env := server.GetEnv()
	request := new(RenderMatrixReq).createFromQueryUrl(req.URL.Query())
	slog.Info("Render matrix", "request", request)

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
		// log.Printf("error marshal response: %+v", err)
		slog.Error("Error marshal response", "msg", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)
}

func GetNextStep(w http.ResponseWriter, req *http.Request) {
	var req_body GetNextStepReq

	slog.Info("Next Step")
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
		// log.Printf("error marshal response: %+v", err)
		slog.Error("Error marshal response", "msg", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)
}

func Healtcheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
