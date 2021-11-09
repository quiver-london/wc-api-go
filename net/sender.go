package net // import "github.com/quiver-london/wc-api-go/v3/net"

import (
	"bytes"
	"encoding/json"
	"github.com/quiver-london/wc-api-go/v3/request"
	"io"
	"net/http"
)

// Sender provides HTTP Requests
type Sender struct {
	requestEnricher RequestEnricher
	urlBuilder      URLBuilder
	httpClient      Client
	requestCreator  RequestCreator
}

// Send method sends requests to WooCommerce API
func (s *Sender) Send(req request.Request) (resp *http.Response, err error) {
	request := s.prepareRequest(req)
	return s.httpClient.Do(request)
}

func (s *Sender) prepareRequest(req request.Request) *http.Request {
	URL := s.urlBuilder.GetURL(req)

	var body io.Reader
	if req.Values != nil && ("POST" == req.Method || "PUT" == req.Method) {
		bodyBytes, err := json.Marshal(req.Values)
		if err != nil {
			panic(err)
		}
		body = bytes.NewReader(bodyBytes)
	}
	request, _ := s.requestCreator.NewRequest(req.Method, URL, body)
	s.requestEnricher.EnrichRequest(request, URL)
	return request
}

// SetRequestEnricher ...
func (s *Sender) SetRequestEnricher(a RequestEnricher) {
	s.requestEnricher = a
}

// SetURLBuilder ...
func (s *Sender) SetURLBuilder(urlBuilder URLBuilder) {
	s.urlBuilder = urlBuilder
}

// SetHTTPClient ...
func (s *Sender) SetHTTPClient(c Client) {
	s.httpClient = c
}

// SetRequestCreator ...
func (s *Sender) SetRequestCreator(rc RequestCreator) {
	s.requestCreator = rc
}
