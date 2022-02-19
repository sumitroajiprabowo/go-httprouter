package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	user := "gordon"
	pass := "secret!"

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/product/:id", GetProuctId)
	router.GET("/product/:id/items/:itemId", GetProuctIdWithItemId)
	router.GET("/image/:file", GetImageinDirectory)
	router.GET("/panic", PanicExample)
	router.POST("/register", Register)
	router.GET("/protected", BasicAuth(Protected, user, pass))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()

}
