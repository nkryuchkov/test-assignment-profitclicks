package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) writeError(w http.ResponseWriter, text string, statusCode int) {
	w.WriteHeader(statusCode)

	e := struct {
		Error string `json:"error"`
	}{
		Error: text,
	}

	data, err := json.Marshal(e)
	if err != nil {
		api.log.Errorf("Could not marshal JSON: %v", err)
		return
	}

	if _, err = w.Write([]byte(data)); err != nil {
		api.log.Errorf("Could not write error: %v", err)
	}
}
