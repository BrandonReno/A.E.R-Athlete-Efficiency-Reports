package data

import(
	"io"
	"encoding/json"
)

func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i) // Create a new encoder and encode the current Workout_Feed to json. Returns an error just in case
}

func FromJSON(i interface{}, r io.Reader) error{
	return json.NewDecoder(r).Decode(i) // Create a new decoder and decode the request body to json. Returns an error just in case
}