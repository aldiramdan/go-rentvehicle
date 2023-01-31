package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status      string      `json:"status"`
	ErrorCode   int         `json:"error-code"`
	RespMessage string      `json:"message"`
	RespData    interface{} `json:"data"`
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) error {

	var resp response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp.Status = "success"
	resp.ErrorCode = code
	resp.RespMessage = getStatus(code)
	resp.RespData = payload

	return json.NewEncoder(w).Encode(resp)

}

func ResponseError(w http.ResponseWriter, code int, message string) error {

	var resp response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp.Status = "error"
	resp.ErrorCode = code
	resp.RespMessage = message

	return json.NewEncoder(w).Encode(resp)
}

func getStatus(code int) (desc string) {
	switch code {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 202:
		desc = "Accepted"
	case 304:
		desc = "Not Modified"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 403:
		desc = "Forbidden"
	case 404:
		desc = "Not Found"
	case 415:
		desc = "Unsupported Media Type"
	case 500:
		desc = "Internal Server Error"
	case 502:
		desc = "Bad Gateway"
	default:
		desc = "Status Code Undefined"
	}
	return
}
