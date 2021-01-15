package api

import (
	v1 "github.com/AhmedARmohamed/FINANCIAL-APP-BACKEND/internal/api/v1"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter provides a handler API service
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/version", v1.VersionHandler)

	return router, nil
}