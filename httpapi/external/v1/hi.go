package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

type hiResponse struct {
	Value string `json:"val"`
}

type PlaceHolder interface {
	Placeholder(ctx context.Context) *http.Response
}

func hiHandler(pc PlaceHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := "bye"
		if resp := pc.Placeholder(r.Context()); resp != nil {
			if resp.StatusCode == http.StatusOK {
				val = "hi from pc client"
			}
		}
		resp := hiResponse{Value: val}
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
