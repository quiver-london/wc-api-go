package client // import "github.com/quiver-london/wc-api-go/v3/client"

import (
	"github.com/quiver-london/wc-api-go/v3/request"
	"net/http"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
