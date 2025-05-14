package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "invalid import path", http.StatusBadRequest)
		return
	}

	user := parts[0]
	repo := parts[1]
	domain := host

	fmt.Fprintf(w, `<html><head>
<meta name="go-import" content="%s/%s/%s git https://%s/%s/%s.git">
</head><body>go get https://%s/%s/%s</body></html>`,
		domain, user, repo,
		domain, user, repo,
		domain, user, repo)

	//fmt.Printf("Served %s/%s for %s\n", user, repo, r.RemoteAddr)
}

func main() {
	port := ":9905"
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
