package rester

import (
	"net/http"

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

//Session is the wrapper of napping.Session.
type Session interface {
	Send(r *napping.Request) (*napping.Response, error)
}

//Rester is a instance to keep http info.
type Rester struct {
	name   string
	apiURL string

	payload interface{}
	err     interface{}

	session Session
}

func (r *Rester) defaultRequest() *Request {
	return &Request{
		Request: napping.Request{
			Header: &http.Header{},
			Error:  r.err,
		},
		session: r.session,
	}
}

// Create will return a new post request.
func (r *Rester) Create(payload interface{}) *Request {
	req := r.defaultRequest()
	req.Method = "POST"
	req.Url = r.apiURL
	req.Payload = payload
	return req
}

// Retrieve object with the id.
func (r *Rester) Retrieve(id string) *Request {
	req := r.defaultRequest()
	req.Method = "GET"
	req.Url = r.apiURL + "/" + id
	return req
}

// List all object of special path, notice that / is the relative root and
// empty is the root path. Howerver, this request should return an array.
func (r *Rester) List(path string) *Request {
	req := r.defaultRequest()
	req.Method = "GET"
	req.Url = r.apiURL + path
	return req
}

// Update something of the id with the payload
func (r *Rester) Update(id string, payload interface{}) *Request {
	req := r.defaultRequest()
	req.Method = "PUT"
	req.Url = r.apiURL + "/" + id
	req.Payload = payload
	return req
}

// Delete object with the id.
func (r *Rester) Delete(id string) *Request {
	req := r.defaultRequest()
	req.Method = "DELETE"
	req.Url = r.apiURL + "/" + id
	return req
}

// Remove objects under path. Notice that "" and "/" is not the same Url.
func (r *Rester) Remove(path string) *Request {
	req := r.defaultRequest()
	req.Method = "DELETE"
	req.Url = r.apiURL + path
	return req
}
