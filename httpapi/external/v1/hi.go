package v1

import (
	"encoding/json"
	"net/http"
)

type hiResponse struct {
	Value string `json:"val"`
}

func hiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := hiResponse{Value: "hi"}
		writeResponse(w, resp)
	}
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enCoder := json.NewEncoder(w)
	err := enCoder.Encode(resp)
	if err != nil {
		return
	}
}
