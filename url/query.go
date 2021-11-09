package url // import "github.com/quiver-london/wc-api-go/v3/url"

import (
	"github.com/quiver-london/wc-api-go/v3/request"
	"net/url"
)

// QueryEnricher uses package auth to enrich existing query parameters with Authentication Based ones
type QueryEnricher interface {
	GetEnrichedQuery(url string, query url.Values, req request.Request) url.Values
}
