package http

import (
	"encoding/json"
	"net/http"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type tasksHandler struct {
	encoder  *encoder
	validate *validator.Validate
}

func newTasksHandler(encoder *encoder) *tasksHandler {
	return &tasksHandler{
		encoder:  encoder,
		validate: validator.New(),
	}
}

func (h *tasksHandler) Routes(r chi.Router) {
	r.Get("/", h.getTasksInfoHandler)
	r.Post("/", h.newTaskHandler)
}

func (h *tasksHandler) getTasksInfoHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("not implemented yet"))
}

func (h *tasksHandler) newTaskHandler(writer http.ResponseWriter, request *http.Request) {
	var task model.Task

	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		h.encoder.Error(writer, err, http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(task); err != nil {
		h.encoder.Error(writer, err, http.StatusBadRequest)
		return
	}

	writer.Write([]byte("not implemented yet"))
}
