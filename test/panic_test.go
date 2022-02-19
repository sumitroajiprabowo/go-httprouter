package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/go-httprouter/exception"
)

func TestPanicHandler(t *testing.T) {

	//router
	router := httprouter.New()

	router.PanicHandler = exception.PanicHandler

	// Add the route.
	router.GET("/panic", PanicExample)

	//request
	request := httptest.NewRequest("GET", "/panic", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Panic : Ups"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestPanicHandler")

}
