package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) json(w http.ResponseWriter, v interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	data, err := json.Marshal(v)
	if err != nil {
		api.log.Errorf("Could not marshal JSON: %v", err)
		api.error(w, "could not marshal JSON", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write([]byte(data)); err != nil {
		api.log.Errorf("Could not write error: %v", err)
		api.error(w, "could not marshal JSON", http.StatusInternalServerError)
	}
}
