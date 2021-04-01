package auth

import (
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/uuid"
)

func VLessGet(id uuid.UUID) *protocol.MemoryUser {
	return authenticator.VLessGet(id)
}
