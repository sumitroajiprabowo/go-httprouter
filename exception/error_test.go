package exception

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {

	//router
	router := httprouter.New()

	// Add the route.
	router.PanicHandler = PanicHandler

	//request
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))

	//response
	response := recorder.Result()

	//read the response body
	resultBody, _ := ioutil.ReadAll(response.Body)

	//expected
	expected := "Panic : Ups"

	//assert
	assert.Equal(t, expected, string(resultBody))

	fmt.Println(recorder.Body.String())

	fmt.Println("Success TestPanicHandler")

}
