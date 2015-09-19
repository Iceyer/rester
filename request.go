package rester

import (
	"net/url"

	"github.com/jmcvetta/napping"
)

//Request implement a simple interface for rest request
type Request struct {
	napping.Request
	session Session
	Error   interface{}
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

//WithResponeAs set the body struct
func (r *Request) WithResponeAs(result interface{}) *Request {
	r.Result = result
	return r
}

//Go send the request
func (r *Request) Go() (*napping.Response, error) {
	rsp, _ := r.session.Send(&r.Request)
	return rsp, nil
}
