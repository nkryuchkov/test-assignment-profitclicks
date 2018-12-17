package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nkryuchkov/test-assignment-profitclicks/logger"
	"github.com/nkryuchkov/test-assignment-profitclicks/service"
)

const (
	v1Prefix = "/api/v1"
)

// Config represents an API configuration.
type Config struct {
	Port int `json:"port"`
}

// API represents a REST API server instance.
type API struct {
	config  *Config
	log     *logger.Logger
	service *service.Service
}

// New returns a new API instance.
func New(config *Config, log *logger.Logger, service *service.Service) *API {
	api := &API{
		config:  config,
		log:     log,
		service: service,
	}

	return api
}

// Start starts the API server.
func (api *API) Start() error {
	api.log.Infof("Starting API")

	r := mux.NewRouter()
	s := r.PathPrefix(v1Prefix).Subrouter()

	s.HandleFunc("/number", api.addNumberToList).Methods("POST")
	s.HandleFunc("/list", api.addNumberList).Methods("POST")
	s.HandleFunc("/list", api.deleteNumberList).Methods("DELETE")
	s.HandleFunc("/operation", api.addOperationToList).Methods("POST")
	s.HandleFunc("/operation", api.addNumberList).Methods("GET")

	return http.ListenAndServe(":"+strconv.Itoa(api.config.Port), r)
}
