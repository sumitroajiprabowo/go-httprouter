package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {

	// Create a new request to pass to the handler.
	router := httprouter.New()

	// Add the route.
	router.GET("/", Index)

	//request to the router
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
	expected := "Welcome!\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	/*
		You can also test the handler directly.
		For example, you can test the handler in a unit test.
	*/
	// //check
	// if recorder.Body.String() != "Welcome!\n" {
	// 	t.Error("Expected \"Welcome!\", got \"" + recorder.Body.String() + "\"")
	// }

	// //check
	// if recorder.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
	// 	t.Error("Expected \"text/plain; charset=utf-8\", got \"" + recorder.Header().Get("Content-Type") + "\"")
	// }

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestRouter")
}
