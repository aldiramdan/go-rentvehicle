package libs

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

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	IsError     bool        `json:"isError"`
	Data        interface{} `json:"data,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")

	if res.IsError {
		w.WriteHeader(res.Code)
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func GetResponse(data interface{}, code int, isError bool) *Response {

	if isError {
		return &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}

	}
	return &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
