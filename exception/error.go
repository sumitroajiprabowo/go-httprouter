package exception

import (
	"fmt"
	"net/http"
)

// Create a new request to pass to the handler with panic
func PanicHandler(w http.ResponseWriter, r *http.Request,
	error interface{}) {
	fmt.Fprint(w, "Panic : ", error)
}
