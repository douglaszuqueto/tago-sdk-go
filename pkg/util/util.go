package util

import (
	"time"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
)

// GeneratePayload GeneratePayload
func GeneratePayload() types.Data {
	return types.Data{
		Variable: "temperature",
		Value:    25.5,
		Time:     time.Now(),
	}
}
