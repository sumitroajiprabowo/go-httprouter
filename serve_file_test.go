package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources/files
var resourcesFile embed.FS

func TestServeFile(t *testing.T) {

	//router
	router := httprouter.New()

	// directory
	directory, _ := fs.Sub(resourcesFile, "resources/files")

	// Add the route.
	router.ServeFiles("/files/*filepath", http.FS(directory))

	//request
	request := httptest.NewRequest("GET", "/files/hello.txt", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Hello World\n"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestServeFile")

}
