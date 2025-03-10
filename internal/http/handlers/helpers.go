package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	erro "inspektor/pkg/errors"
)

func (s service) respond(w http.ResponseWriter, data interface{}, status int) {
	var respData interface{}
	switch v := data.(type) {
	case nil:
	case erro.ErrArgument:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
	default:
		respData = data
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(respData)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}

// it reads to the memory.
func (s service) readRequestBody(r *http.Request) ([]byte, error) {
	// Read the content
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			err := errors.New("could not read request body")
			return nil, err
		}
	}
	return bodyBytes, nil
}

// will place the body bytes back to the request body which could be read in subsequent calls on Handlers
// for example, you have more than 1 middleware and each of them need to read the body. If the first middleware read the body
// the second one won't be able to read it, unless you put the request body back.
func (s service) restoreRequestBody(r *http.Request, bodyBytes []byte) {
	// Restore the io.ReadCloser to its original state
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
}
