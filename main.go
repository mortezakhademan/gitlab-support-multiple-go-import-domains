package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) < 2 {
		http.Error(w, "invalid import path", http.StatusBadRequest)
		return
	}

	modulePath := fmt.Sprintf("%s/%s", host, path)
	repoURL := fmt.Sprintf("https://%s/%s.git", host, path)

	fmt.Fprintf(w, `<html><head>
<meta name="go-import" content="%s git %s">
</head><body>go get %s</body></html>`,
		modulePath, repoURL, modulePath)
}

func main() {
	port := ":9905"
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
