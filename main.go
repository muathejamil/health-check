package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "health")
}

func server(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/server", server)

	http.ListenAndServe(":8090", nil)
}
