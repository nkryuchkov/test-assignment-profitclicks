package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) addNumberToList(w http.ResponseWriter, r *http.Request) {

}

func (api *API) addNumberList(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	uid, err := api.service.AddNumberList()
	if err != nil {
		api.log.Errorf("Could not add number list: %v", err)
		api.writeError(w, "could not add number list")
		return
	}

	resp := struct {
		UID string `json:"uid"`
	}{
		UID: uid,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		api.log.Errorf("Could not marshal JSON: %v", err)
		api.writeError(w, "could not marshal JSON")
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(data); err != nil {
		api.log.Errorf("Could not write response: %v", err)
		return
	}
}

func (api *API) deleteNumberList(w http.ResponseWriter, r *http.Request) {

}

func (api *API) addOperationToList(w http.ResponseWriter, r *http.Request) {

}

func (api *API) getListResult(w http.ResponseWriter, r *http.Request) {

}

func (api *API) setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
