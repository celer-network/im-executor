package server

import (
	"encoding/json"
	"net/http"

	"github.com/celer-network/im-executor/server/models"
)

func handler[T any, R any](handleRequest func(*R) (T, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &models.ServerResponse[T]{
			Code: http.StatusOK,
		}
		if r.Method != http.MethodPost {
			response.Code = http.StatusMethodNotAllowed
			response.Message = "only 'POST' is allowed"
			bz, _ := json.Marshal(response)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(bz)
			return
		}
		body := new(R)
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(body)
		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "cannot parse request body: " + err.Error()
			bz, _ := json.Marshal(response)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(bz)
			return
		}
		res, err := handleRequest(body)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
		}
		response.Data = res
		response.Message = "success"
		bz, _ := json.Marshal(response)
		w.Write(bz)
	}
}
