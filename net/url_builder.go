package net // import "github.com/quiver-london/wc-api-go/v3/net"

import (
	"github.com/quiver-london/wc-api-go/v3/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
