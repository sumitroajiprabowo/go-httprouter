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

func TestHandlerNotFound(t *testing.T) {

	//router
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(exception.NotFoundHandler)

	//request
	request := httptest.NewRequest("GET", "/", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Not Found"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestHandlerNotFound")

}
