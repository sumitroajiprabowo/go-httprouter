package middleware

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func LoggingMiddlerware(handler http.Handler) *LogMiddleware {
	return &LogMiddleware{Handler: handler}
}

func (lm *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log from middleware")
	lm.Handler.ServeHTTP(w, r)
}
