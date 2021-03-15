package auth

import (
	"github.com/xtls/xray-core/common/protocol"
)

func TrojanGet(hash string) *protocol.MemoryUser {
	return authenticator.TrojanGet(hash)
}