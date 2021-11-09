package url

import (
	"github.com/quiver-london/wc-api-go/v3/request"
	URL "net/url"
)

// QueryEnricherMock ...
type QueryEnricherMock struct {
	query URL.Values
}

// GetEnrichedQuery ...
func (q *QueryEnricherMock) GetEnrichedQuery(url string, query URL.Values, req request.Request) URL.Values {
	return q.query
}
