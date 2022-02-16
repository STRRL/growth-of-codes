package api

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "growth-of.codes API under construction")
}
