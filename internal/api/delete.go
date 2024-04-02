package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const uploadDirectory = "./uploads/"

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	if filename == "" {
		RespondWithErr(w, http.StatusBadRequest, fmt.Errorf("ilename parameter is missing"))
		return
	}

	filePath := filepath.Join(uploadDirectory, filename)
	err := os.Remove(filePath)
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, fmt.Errorf("error deleting file %w", err))
		return
	}

	fmt.Fprintf(w, "File deleted successfully: %s", filename)
}
