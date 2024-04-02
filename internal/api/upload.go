package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		RespondWithErr(w, http.StatusBadRequest, fmt.Errorf("error retrieving file %w", err))
		return
	}
	defer file.Close()

	if _, err := os.Stat(uploadDirectory); os.IsNotExist(err) {
		os.Mkdir(uploadDirectory, 0755)
	}

	createdFile, err := os.OpenFile(uploadDirectory+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, fmt.Errorf("error saving file %w", err))
		return
	}
	defer createdFile.Close()

	_, err = io.Copy(createdFile, file)
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, fmt.Errorf("error saving file %w", err))
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}
