package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternNamedParams(t *testing.T) {

	// Create a new request to pass to the handler.
	router := httprouter.New()

	// Add the route.
	router.GET("/product/:id/items/:itemId", GetProuctIdWithItemId)

	//request to the router
	request := httptest.NewRequest("GET", "/product/123/items/456", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "ID 123 with Item ID: 456"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestRouterPatternNamedParams")

}

func TestRouterPatternCatchAllParams(t *testing.T) {

	// Create a new request to pass to the handler.
	router := httprouter.New()

	// Add the route.
	router.GET("/images/*image", GetImageinDirectory)

	//request to the router
	request := httptest.NewRequest("GET", "/images/thumbnail/profile.png", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Image /thumbnail/profile.png"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestRouterPatternCatchAllParams")

}
