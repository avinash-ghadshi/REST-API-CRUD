package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  bool
	Message string
}

func SendResponse(w http.ResponseWriter, code int, msg string, status bool) bool {
	w.WriteHeader(code)
	var resp Response
	resp.Status = status
	resp.Message = msg
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
