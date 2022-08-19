package url // import "github.com/quiver-london/wc-api-go/v3/url"

import (
	"fmt"
	URL "net/url"
	"strings"

	"github.com/quiver-london/wc-api-go/v3/options"
	"github.com/quiver-london/wc-api-go/v3/request"
)

// Builder structure
type Builder struct {
	queryEnricher QueryEnricher
	options       options.Basic
}

// GetURL method prepare URL be adding required authentication parameter values
func (b *Builder) GetURL(req request.Request) string {
	query := b.getFilteredQuery(req)
	urlWithEndpoint := b.getBaseURL() + req.Endpoint
	if query.Encode() != "" {
		urlWithEndpoint += fmt.Sprintf("?%s", query.Encode())
	}
	return urlWithEndpoint
}

func (b *Builder) getFilteredQuery(req request.Request) URL.Values {
	query := URL.Values{}
	if req.Method == "GET" || req.Method == "DELETE" {
		for k, v := range req.Values {
			query.Add(k, v.(string))
		}
	}
	return query
}

// GetBaseURL method prepare BaseURL according to Options
func (b *Builder) getBaseURL() string {
	return strings.TrimRight(b.options.URL, "/") + b.getAPIPrefix() + b.options.Version() + "/"
}

func (b *Builder) getAPIPrefix() string {
	if b.options.WPAPI() {
		return b.options.WPAPIPrefix()
	}
	return options.DefaultAPIPrefix
}

// SetOptions method sets WooCommerce integration options to structure's inner variable
func (b *Builder) SetOptions(o options.Basic) {
	b.options = o
}

// SetQueryEnricher ...
func (b *Builder) SetQueryEnricher(qe QueryEnricher) {
	b.queryEnricher = qe
}
