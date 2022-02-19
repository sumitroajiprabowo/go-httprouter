package exception

import (
	"fmt"
	"net/http"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method Not Allowed")
}
