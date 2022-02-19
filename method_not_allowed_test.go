package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/go-httprouter/exception"
)

func TestMethodNotAllowed(t *testing.T) {

	//router
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowedHandler)

	// Add the route.
	router.POST("/register", Register)

	//request
	request := httptest.NewRequest("PUT", "/register", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Method Not Allowed"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestMethodNotAllowed")
}
