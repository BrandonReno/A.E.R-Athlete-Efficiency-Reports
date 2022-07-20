package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	mocks "github.com/BrandonReno/A.E.R/mocks/gen"
	"github.com/BrandonReno/A.E.R/server"
	"github.com/go-chi/chi"
)

type BackendFixture struct {
	Router      chi.Router
	TestServer  *httptest.Server
	AthleteRepo *mocks.AthleteRepository
	WorkoutRepo *mocks.WorkoutRepository
}

func NewBackendFixture() *BackendFixture {
	return &BackendFixture{
		Router:      chi.NewRouter(),
		AthleteRepo: new(mocks.AthleteRepository),
		WorkoutRepo: new(mocks.WorkoutRepository),
	}
}

func (fx *BackendFixture) MakeRequest(method, url string, payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func (fx *BackendFixture) UnmarshallResponseData(resp *http.Response, target interface{}) error {
	jsonResp := new(server.JSONResponse)
	if err := json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
		return err
	}
	b, err := json.Marshal(jsonResp.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}

func (fx *BackendFixture) MakeURL(path string, parameters ...interface{}) string {
	return MakeURLFromServer(fx.TestServer, path, parameters...)
}

func MakeURLFromServer(server *httptest.Server, path string, parameters ...interface{}) string {
	return fmt.Sprintf("%s%s", server.URL, fmt.Sprintf(path, parameters...))
}
