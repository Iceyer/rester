package rester

import (
	"fmt"
	"net/http"
)

//Rest create an instance  to send/recive data from rest server.
func Rest() *Rester {
	fmt.Print("a")
	return &Rester{}
}

//Rester is a instance to keep http info.
type Rester struct {
	header http.Header
}

//AddHeader will add a header params to request.
func (r *Rester) AddHeader(key, value string) *Rester {
	r.header.Add(key, value)
	return r
}
