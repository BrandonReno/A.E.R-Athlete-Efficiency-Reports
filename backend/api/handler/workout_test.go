package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrandonReno/A.E.R/mocks"
	"github.com/BrandonReno/A.E.R/models"
	mockery "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWorkout_GetAll(t *testing.T) {
	t.Run("GetAll -- Success", func(t *testing.T) {
		SetupEndpoints(t, func(fixture *mocks.BackendFixture) {
			fixture.WorkoutRepo.On("GetAll", mockery.Anything).Return([]*models.Workout{{Workout_ID: 1}}, nil)
			resp, err := fixture.MakeRequest(http.MethodGet, fixture.MakeURL("/workouts"), nil)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, http.StatusOK)
			var workouts []*models.Workout
			err = fixture.UnmarshallResponseData(resp, &workouts)
			require.NoError(t, err)
			require.NotNil(t, workouts)
			require.Equal(t, workouts[0].Workout_ID, 1)
		})
	})
}

func TestWorkout_GetOne(t *testing.T) {
	t.Run("GetOne -- Success", func(t *testing.T) {
		SetupEndpoints(t, func(fixture *mocks.BackendFixture) {
			workoutID := 1
			fixture.WorkoutRepo.On("GetByID", mockery.Anything, workoutID).Return(models.Workout{Workout_ID: workoutID}, nil)
			resp, err := fixture.MakeRequest(http.MethodGet, fixture.MakeURL("/workouts/1"), nil)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, http.StatusOK)
			var workout *models.Workout
			err = fixture.UnmarshallResponseData(resp, workout)
			require.NoError(t, err)
			require.NotNil(t, workout)
			require.Equal(t, workout.Workout_ID, workoutID)
		})
	})
}

func SetupEndpoints(t *testing.T, testBody func(fixture *mocks.BackendFixture)) {
	fixture := mocks.NewBackendFixture()
	handler := NewWorkoutHandler(fixture.WorkoutRepo)
	handler.MountRoutes(fixture.Router)
	fixture.TestServer = httptest.NewServer(fixture.Router)
	defer fixture.TestServer.Close()
	testBody(fixture)
}
