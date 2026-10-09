package jsonresponder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	JSONContentType = "application/json"
	JSONCharset     = "utf-8"
)

type JSONResponder interface {
	Write(w http.ResponseWriter, statusCode int, data interface{}) error
	WriteSimple(w http.ResponseWriter, statusCode int, message string) error
}

type jsonResponder struct {
	contentType string
}

func NewDefaultJSONResponder() JSONResponder {
	return &jsonResponder{
		contentType: fmt.Sprintf("%s; charset=%s", JSONContentType, JSONCharset),
	}
}

func (j *jsonResponder) Write(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", j.contentType)
	w.WriteHeader(statusCode)

	if data == nil {
		return nil
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	_, err = w.Write(content)
	return err
}

func (j *jsonResponder) WriteSimple(w http.ResponseWriter, statusCode int, message string) error {
	return j.Write(w, statusCode, SimpleMessage{Message: message})
}
