package exception

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not Found")
}
