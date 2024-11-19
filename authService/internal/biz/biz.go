package biz

import (
	"fmt"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet()

func GenerateKey(userID int32) string {
	return fmt.Sprintf("auth:%d", userID)
}
