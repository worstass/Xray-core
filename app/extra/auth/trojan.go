package auth

import (
	"github.com/xtls/xray-core/common/protocol"
)

func TrojanGet(hash string) *protocol.MemoryUser {
	for _, a := range authenticators {
		res := a.TrojanGet(hash)
		if res != nil {
			return res
		}
	}
	return nil
}