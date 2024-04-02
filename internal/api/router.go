package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{router: mux.NewRouter()}
}

func (r *Router) SetupRoutes() {
	r.router.HandleFunc("/", HandleRoot).Methods("GET")
	r.router.HandleFunc("/upload", HandleUpload).Methods("POST")
	r.router.HandleFunc("/download", HandleDownload).Methods("GET")
	r.router.HandleFunc("/delete/{filename}", HandleDelete).Methods("DELETE")
	r.router.HandleFunc("/rename", HandleRename).Methods("PUT")
}

func (r *Router) StartServer() error {
	return http.ListenAndServe("127.0.0.1:8080", r.router)
}
