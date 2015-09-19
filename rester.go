package rester

import (
	"net/http"
	"net/url"

	"github.com/jmcvetta/napping"
)

//Rest create an instance  to send/recive data from rest server.
func Rest(name string, url string) *Rester {
	return &Rester{
		name:    name,
		apiURL:  url,
		session: &napping.Session{},
	}
}

//Session is the wrapper of napping.Session
type Session interface {
	Send(r *napping.Request) (*napping.Response, error)
}

//Rester is a instance to keep http info.
type Rester struct {
	name string

	apiURL  string
	headers http.Header
	params  url.Values

	payload interface{}
	err     interface{}

	session Session
}

//Request implement a simple interface for rest request
type Request struct {
	napping.Request
	session Session
	err     interface{}
}

// Create will return a new post request
func (r *Rester) Create(payload interface{}) *Request {
	return &Request{
		Request: napping.Request{
			Url:     r.apiURL,
			Method:  "POST",
			Payload: r.payload,
			Header:  &http.Header{},
		},
		session: r.session,
	}
}

//AddHeader will add a header params to request.
func (r *Request) AddHeader(key, value string) *Request {
	r.Header.Add(key, value)
	return r
}

//WithParams add query string params
func (r *Request) WithParams(p *url.Values) *Request {
	r.Params = p
	return r
}

//WithPayloadAs set the body struct
func (r *Request) WithPayloadAs(payload interface{}) *Request {
	r.Result = payload
	return r
}

//Go send the request
func (r *Request) Go() (*napping.Response, error) {
	rsp, _ := r.session.Send(&r.Request)
	return rsp, nil
}
