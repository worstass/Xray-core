package auth

import "github.com/xtls/xray-core/common/protocol"

func VMessGetUser(email string) *protocol.MemoryUser {
	return authenticator.VMessGetUser(email)
}
