package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogMiddleware(t *testing.T) {

	//handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	//request
	request := httptest.NewRequest("GET", "/", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	LoggingMiddlerware(handler).ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Hello World"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestLogMiddleware")

}
