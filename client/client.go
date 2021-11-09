package client // import "github.com/quiver-london/wc-api-go/v3/client"

import (
	"github.com/quiver-london/wc-api-go/v3/request"
	"net/http"
	"net/url"
)

// Client is upper level class which delegate all work to Requester
type Client struct {
	sender Sender
}

// Get Method loads data from Endpoint with specified parameters
func (c *Client) Get(endpoint string, parameters url.Values) (*http.Response, error) {
	return c.sender.Send(request.Request{
		Method:   "GET",
		Endpoint: endpoint,
		Values:   parameters,
	})
}

// Post Method usually creates new instances
func (c *Client) Post(endpoint string, data url.Values) (*http.Response, error) {
	return c.sender.Send(request.Request{
		Method:   "POST",
		Endpoint: endpoint,
		Values:   data,
	})
}

// Put Method usually update existing instances
func (c *Client) Put(endpoint string, data url.Values) (*http.Response, error) {
	return c.sender.Send(request.Request{
		Method:   "PUT",
		Endpoint: endpoint,
		Values:   data,
	})
}

// Delete Method usually removes existing instances
func (c *Client) Delete(endpoint string, parameters url.Values) (*http.Response, error) {
	return c.sender.Send(request.Request{
		Method:   "DELETE",
		Endpoint: endpoint,
		Values:   parameters,
	})
}

// Options Method usually using for checking possibility of POST requests
func (c *Client) Options(endpoint string) (*http.Response, error) {
	return c.sender.Send(request.Request{
		Method:   "OPTIONS",
		Endpoint: endpoint,
		Values:   nil,
	})
}
