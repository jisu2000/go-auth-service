package wrapper

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, success bool, data interface{}, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := ApiResponse{
		Success: success,
		Data:    data,
		Error:   errMsg,
	}
	_ = json.NewEncoder(w).Encode(res)
}
