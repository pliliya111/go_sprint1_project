package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pliliya111/go_sprint1_project/internal/handler"
)

func StartServer(timeout time.Duration) {

	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	fmt.Println("Starting server on :8080...")

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	StartServer(0)
}
