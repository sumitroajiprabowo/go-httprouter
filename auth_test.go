package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {

	//router
	router := httprouter.New()

	// Add the route.
	router.GET("/", BasicAuth(func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	}, "user", "password"))

	//request
	request := httptest.NewRequest("GET", "/", nil)

	//request header with basic auth
	request.SetBasicAuth("user", "password")

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Hello World"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(string(resultBody))

	fmt.Println("Success TestBasicAuth")
}
