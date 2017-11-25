package daemon

import (
	"fmt"
	"net/http"
)

// PongHandler Return pong
func PongHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello")
		},
	)
}