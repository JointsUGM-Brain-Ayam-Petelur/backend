package auth

import (
	"hearable/backend/package/jsonresponder"
	"net/http"
)

type Handler struct {
	service       serviceInterface
	jsonResponder jsonresponder.JSONResponder
}

type serviceInterface interface {
	Test(response string) string
}

func NewHandler() *Handler {
	return &Handler{
		service:       NewService(),
		jsonResponder: jsonresponder.NewDefaultJSONResponder(),
	}
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	response := h.service.Test("Hello, World!")
	err := h.jsonResponder.WriteSimple(w, http.StatusOK, response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
