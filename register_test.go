package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	//router
	router := httprouter.New()

	// Add the route.
	router.POST("/register", Register)

	//request
	request := httptest.NewRequest("POST", "/register", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Register Success!\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(string(resultBody))

	fmt.Println("Success TestRegister")

}
