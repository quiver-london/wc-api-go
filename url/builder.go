package url // import "github.com/quiver-london/wc-api-go/v3/url"

import (
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
	values := b.queryEnricher.GetEnrichedQuery(urlWithEndpoint, query, req)
	return urlWithEndpoint + "?" + values.Encode()
}

func (b *Builder) getFilteredQuery(req request.Request) URL.Values {
	var query URL.Values
	if req.Method == "GET" || req.Method == "DELETE" {
		query := make(URL.Values)
		for k, v := range req.Values {
			query[k] = []string{v.(string)}
		}
	} else {
		query = nil
	}
	if nil == query {
		query = URL.Values{}
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
