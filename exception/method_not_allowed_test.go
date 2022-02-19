package exception

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerMethodNotAllowed(t *testing.T) {

	//handler
	handler := http.HandlerFunc(MethodNotAllowedHandler)

	//request
	request := httptest.NewRequest("PUT", "/register", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	handler.ServeHTTP(recorder, request)

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
