package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ReadJSON(r *http.Request, receiver interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, receiver); err != nil {
		return err
	}
	if validatable, ok := receiver.(validation.Validatable); ok {
		return validatable.Validate()
	}
	return nil
}
