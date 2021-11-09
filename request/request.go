package request // import "github.com/quiver-london/wc-api-go/v3/request"

import (
	"net/url"
)

// Request ...
type Request struct {
	Method   string
	Endpoint string
	Values   url.Values
}
