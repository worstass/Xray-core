package auth

import (
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/uuid"
)

func VLessGet(id uuid.UUID) *protocol.MemoryUser {
	for _, a := range authenticators {
		res := a.VLessGet(id)
		if res != nil {
			return res
		}
	}
	return nil
}
