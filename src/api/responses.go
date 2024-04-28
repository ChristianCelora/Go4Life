package api

import (
	"fmt"
	"strings"
)

type ApiErrorRes struct {
	Code int
	Msg  string
}

type RenderMatrixRes struct {
	Matrix [][]uint8
}

func (res *RenderMatrixRes) MarshalJson() ([]byte, error) {
	var matrix string
	if res.Matrix == nil {
		matrix = "null"
	} else {
		matrix = strings.Join(
			strings.Fields(fmt.Sprintf("%d", res.Matrix)),
			",",
		)
	}
	jsonResult := fmt.Sprintf(`{"matrix":%s}`, matrix)
	return []byte(jsonResult), nil
}
