package http

import (
	"encoding/json"
	"net/http"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type tasksHandler struct {
	encoder     *encoder
	validate    *validator.Validate
	taskService taskService
}

func newTasksHandler(encoder *encoder, taskService taskService) *tasksHandler {
	return &tasksHandler{
		encoder:     encoder,
		validate:    validator.New(),
		taskService: taskService,
	}
}

func (h *tasksHandler) Routes(r chi.Router) {
	r.Get("/", h.getTasksInfoHandler)
	r.Post("/", h.newTaskHandler)
}

func (h *tasksHandler) getTasksInfoHandler(writer http.ResponseWriter, request *http.Request) {
	tasks := h.taskService.GetTasks()

	h.encoder.JSONResponse(writer, tasks)
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

	go h.taskService.AddTask(task)

	h.encoder.StatusOK(writer)
}
