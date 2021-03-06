package api

import (
	"net/http"
	"strconv"

	"github.com/nkryuchkov/test-assignment-profitclicks/service"
)

func (api *API) addNumberToList(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	vars := r.URL.Query()

	uids, ok := vars["uid"]
	if !ok || len(uids) == 0 {
		api.error(w, "no UID specified", http.StatusBadRequest)
		return
	}
	uid := uids[0]

	numbers, ok := vars["number"]
	if !ok || len(numbers) == 0 {
		api.error(w, "no number specified", http.StatusBadRequest)
		return
	}
	numberStr := numbers[0]
	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		api.error(w, "bad number specified", http.StatusBadRequest)
		return
	}

	err = api.service.AddNumberToList(uid, number)
	if err == service.ErrListNotExists {
		api.error(w, "list with the given UID does not exist", http.StatusBadRequest)
		return
	}
	if err != nil {
		api.log.Errorf("Could not add number to list: %v", err)
		api.error(w, "could not add number to list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) addNumberList(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	uid, err := api.service.AddNumberList()
	if err != nil {
		api.log.Errorf("Could not add number list: %v", err)
		api.error(w, "could not add number list", http.StatusInternalServerError)
		return
	}

	resp := struct {
		UID string `json:"uid"`
	}{
		UID: uid,
	}

	api.json(w, resp, http.StatusOK)
}

func (api *API) deleteNumberList(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	vars := r.URL.Query()

	uids, ok := vars["uid"]
	if !ok || len(uids) == 0 {
		api.error(w, "no uid specified", http.StatusBadRequest)
		return
	}
	uid := uids[0]

	if err := api.service.DeleteNumberList(uid); err != nil {
		api.log.Errorf("Could not delete number list: %v", err)
		api.error(w, "could not delete number list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) addOperationToList(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	vars := r.URL.Query()

	uids, ok := vars["uid"]
	if !ok || len(uids) == 0 {
		api.error(w, "no UID specified", http.StatusBadRequest)
		return
	}
	uid := uids[0]

	names, ok := vars["name"]
	if !ok || len(names) == 0 {
		api.error(w, "no name specified", http.StatusBadRequest)
		return
	}
	name := names[0]

	err := api.service.AddOperationToList(uid, name)
	if err == service.ErrListNotExists {
		api.error(w, "list with the given UID does not exist", http.StatusBadRequest)
		return
	}
	if err == service.ErrOperationNotExists {
		api.error(w, "operation with the given name does not exist", http.StatusBadRequest)
		return
	}
	if err != nil {
		api.log.Errorf("Could not add operation to list: %v", err)
		api.error(w, "could not add operation to list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) getListResult(w http.ResponseWriter, r *http.Request) {
	api.setHeaders(w)

	vars := r.URL.Query()

	uids, ok := vars["uid"]
	if !ok || len(uids) == 0 {
		api.error(w, "no UID specified", http.StatusBadRequest)
		return
	}
	uid := uids[0]

	result, err := api.service.GetListResult(uid)
	if err == service.ErrListNotExists {
		api.error(w, "list with the given UID does not exist", http.StatusBadRequest)
		return
	}

	if err == service.ErrNoOperationSet {
		api.error(w, "no operation set for this list", http.StatusBadRequest)
		return
	}

	if err != nil {
		api.log.Errorf("Could not get list result: %v", err)
		api.error(w, "could not get list result", http.StatusInternalServerError)
		return
	}

	resp := struct {
		Result int64 `json:"result"`
	}{
		Result: result,
	}

	api.json(w, resp, http.StatusOK)
}

func (api *API) setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
