package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		logrus.Error("error arose when serving http")
	}
}
