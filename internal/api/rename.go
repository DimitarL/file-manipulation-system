package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func HandleRename(w http.ResponseWriter, r *http.Request) {
	oldName := r.FormValue("oldname")
	newName := r.FormValue("newname")

	if oldName == "" || newName == "" {
		RespondWithErr(w, http.StatusBadRequest, fmt.Errorf("old name or new name parameter is missing"))
		return
	}

	oldPath := filepath.Join(uploadDirectory, oldName)
	newPath := filepath.Join(uploadDirectory, newName)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, fmt.Errorf("error renaming file %w", err))
		return
	}

	fmt.Fprintf(w, "File renamed successfully from %s to %s", oldName, newName)
}
