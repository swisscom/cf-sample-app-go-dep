package hello

import (
	"fmt"
	"net/http"
)

// Handler for /
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!\n\nWelcome to the Swisscom Application Cloud!", r.Header.Get("User-Agent"))
}
