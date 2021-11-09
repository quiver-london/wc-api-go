package auth // import "github.com/quiver-london/wc-api-go/v3/auth"

import (
	"strconv"
	"time"
)

// MicroTimer ...
type MicroTimer struct {
}

// Get current micro time
func (m *MicroTimer) Get() string {
	mc := time.Now().UnixNano()
	return strconv.FormatInt(mc, 10)
}
