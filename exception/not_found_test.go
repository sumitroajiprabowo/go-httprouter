package exception

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerNotFound(t *testing.T) {

	//handler
	handler := http.HandlerFunc(NotFoundHandler)

	//request
	request := httptest.NewRequest("GET", "/", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	handler.ServeHTTP(recorder, request)

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
