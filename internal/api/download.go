package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		RespondWithErr(w, http.StatusBadRequest, fmt.Errorf("filename parameter is missing"))
		return
	}

	filePath := filepath.Join(uploadDirectory, filename)
	file, err := os.Open(filePath)
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, fmt.Errorf("error opening file %w", err))
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error downloading file", http.StatusInternalServerError)
		return
	}
}
