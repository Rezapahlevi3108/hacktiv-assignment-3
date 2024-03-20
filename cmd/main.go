package main

import (
	"net/http"

	"github.com/Rezapahlevi3108/hacktiv-assignment-3/internal/handler"
)

func main() {
	// Mulai auto-reload status setiap 15 detik
	handler.StartAutoReload()

	// Routing HTTP
	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/", handler.IndexHandler)

	// Jalankan server HTTP
	http.ListenAndServe(":8080", nil)
}
