package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BrandonReno/A.E.R/models"
	"github.com/BrandonReno/A.E.R/repository"
	"github.com/BrandonReno/A.E.R/server"
	"github.com/go-chi/chi"
)

type WorkoutHandler struct {
	workoutRepository repository.WorkoutRepository
}

func NewWorkoutHandler(wr repository.WorkoutRepository) *WorkoutHandler {
	return &WorkoutHandler{workoutRepository: wr}
}

func (wr WorkoutHandler) MountRoutes(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Method(http.MethodGet, "/workouts", server.Handler(wr.GetAll))
		r.Method(http.MethodPost, "/workouts", server.Handler(wr.Create))
		r.Method(http.MethodGet, "/workouts", server.Handler(wr.GetOne))
	})
}

func (wr *WorkoutHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	workouts, err := wr.workoutRepository.GetAll(r.Context())
	if err != nil {
		return err
	}
	return server.WriteSuccessResponse(w, http.StatusOK, workouts)
}

func (wr *WorkoutHandler) GetOne(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		server.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no id in request"))
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		server.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("id can not be changed to int"))
	}
	workout, err := wr.workoutRepository.GetByID(r.Context(), intID)
	if err != nil {
		server.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
	return server.WriteSuccessResponse(w, http.StatusOK, workout)
}

func (wr *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) error {
	var workout *models.Workout
	err := server.ReadJSON(r, workout)
	if err != nil {
		server.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
	if err = wr.workoutRepository.Create(r.Context(), workout); err != nil {
		return server.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}
	return server.WriteSuccessResponse(w, http.StatusCreated, workout)
}
