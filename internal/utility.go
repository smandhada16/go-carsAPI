package internal

import (
	"encoding/json"
	"net/http"
)

func returnErrorResponse(code int, msg string, resp http.ResponseWriter) {

	//create error response json
	errRespn := Response{
		Code:    code,
		Message: msg,
	}

	respnseByte, err := json.Marshal(errRespn)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Internal Server Error"))
	}

	//write repsonse code header
	resp.WriteHeader(code)
	resp.Write(respnseByte)
}

func isEmpty(s string) bool {
	if len(s) == 0 || s == "" {
		return true
	}
	return false
}
