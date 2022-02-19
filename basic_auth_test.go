package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthorized(t *testing.T) {

	//router
	router := httprouter.New()

	// Add the route.
	router.GET("/protected", BasicAuth(Protected, "gordon", "secret!"))

	//request
	request := httptest.NewRequest("GET", "/protected", nil)

	//request header with basic auth
	request.SetBasicAuth("gordon", "secret!")

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Protected!\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(string(resultBody))

	fmt.Println("Success TestAuthorized")

}

func TestUnAuthorized(t *testing.T) {

	//router
	router := httprouter.New()

	// Add the route.
	router.GET("/protected", BasicAuth(Protected, "gordon", "secret!"))

	//request
	request := httptest.NewRequest("GET", "/protected", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Unauthorized\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestUnAuthorized")

}

func TestWrongUserNameOrPassword(t *testing.T) {

	//router
	router := httprouter.New()

	// Add the route.
	router.GET("/protected", BasicAuth(Protected, "gordon", "secret!"))

	//request
	request := httptest.NewRequest("GET", "/protected", nil)

	//request header with basic auth
	request.SetBasicAuth("gordon", "wrongpassword")

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Unauthorized\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestWrongUserNameOrPassword")

}
