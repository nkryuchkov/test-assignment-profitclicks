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
		api.writeError(w, "could not add number list", http.StatusInternalServerError)
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
		api.writeError(w, "could not marshal JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(data); err != nil {
		api.log.Errorf("Could not write response: %v", err)
		return
	}
}

func (api *API) deleteNumberList(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	vars := r.URL.Query()

	uids, ok := vars["uid"]
	if !ok || len(uids) == 0 {
		api.writeError(w, "no uid specified", http.StatusBadRequest)
		return
	}
	uid := uids[0]

	if err := api.service.DeleteNumberList(uid); err != nil {
		api.log.Errorf("Could not delete number list: %v", err)
		api.writeError(w, "could not delete number list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) addOperationToList(w http.ResponseWriter, r *http.Request) {

}

func (api *API) getListResult(w http.ResponseWriter, r *http.Request) {

}

func (api *API) setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
