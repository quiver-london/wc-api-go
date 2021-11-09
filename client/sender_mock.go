package client

import (
	"github.com/quiver-london/wc-api-go/v3/request"
	"net/http"
)

// SenderMock imitates sending requests and receiving responses
type SenderMock struct {
	response http.Response
}

// Send ...
func (r *SenderMock) Send(req request.Request) (resp *http.Response, err error) {
	return &r.response, nil
}
