package api

import (
	"net/url"
	"strconv"
)

const (
	MAX_MATRIX_ROWS = 100
	MAX_MATRIX_COLS = 100
)

type RenderMatrixReq struct {
	template string
	offsetX  int
	offsetY  int
	rows     int
	cols     int
}

func (request RenderMatrixReq) createFromQueryUrl(req_query url.Values) RenderMatrixReq {
	request.template = req_query.Get("template")
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

	return request
}

type GetNextStepReq struct {
	Matrix [][]uint8 `json:"matrix"`
}
