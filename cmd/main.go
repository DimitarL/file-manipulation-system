package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DimitarL/file-manipulation-system/internal/api"
)

func main() {
	srv := api.NewRouter()
	srv.SetupRoutes()

	fmt.Println("Server started on 127.0.0.1:8080")
	err := srv.StartServer()
	if err != nil {
		slog.With(slog.String("error", err.Error())).Error("failed to listen and serve")
		os.Exit(1)
	}
}
