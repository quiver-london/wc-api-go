package request // import "github.com/quiver-london/wc-api-go/v3/request"

// Request ...
type Request struct {
	Method   string
	Endpoint string
	Values   map[string]interface{}
}
