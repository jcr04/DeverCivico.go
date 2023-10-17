package presentation

import (
	"encoding/json"
	"net/http"
)

func decodeJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
