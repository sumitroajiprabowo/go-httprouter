package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Create a new request to pass to the handler.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Create a new request to pass to the handler with params.
func GetProuctId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "ID: %s\n", ps.ByName("id"))
}

/*
Create a new request to pass to the handler with router pattern named params.

Pattern: /user/:user

 /user/gordon              match
 /user/you                 match
 /user/gordon/profile      no match
 /user/

Reference : https://github.com/julienschmidt/httprouter
*/
func GetProuctIdWithItemId(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params) {

	fmt.Fprintf(w, "ID %s with Item ID: %s", ps.ByName("id"), ps.ByName("itemId"))

}

/*
Create a new request to pass to the handler with router pattern catch all params.

The second type are catch-all parameters and have the form *name. Like the name suggests, they match everything. Therefore they must always be at the end of the pattern:

Pattern: /src/*filepath

 /src/                     match
 /src/somefile.go          match
 /src/subdir/somefile.go   match

Reference : https://github.com/julienschmidt/httprouter
*/
func GetImageinDirectory(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params) {
	fmt.Fprintf(w, "Image %s", ps.ByName("image"))
}

// Create a new request to pass to the handler with panic
func PanicExample(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params) {
	panic("Ups")
}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Register Success!\n")
}

func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}
