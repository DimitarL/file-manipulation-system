package api

import (
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}
