package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/go-httprouter/middleware"
)

func TestMiddleware(t *testing.T) {

	//router
	router := httprouter.New()

	//add router
	router.GET("/", Index)

	//add middleware
	middleware := middleware.LogMiddleware{
		Handler: router,
	}

	//request
	request := httptest.NewRequest("GET", "/", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	middleware.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Welcome!\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestMiddleware")

}
