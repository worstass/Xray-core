package auth

import (
	"crypto/cipher"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/uuid"
)

type Authenticator interface {
	VLessGet(id uuid.UUID) *protocol.MemoryUser
	TrojanGet(hash string) *protocol.MemoryUser
	VMessTimedUserValidatorGet(userHash []byte) (*protocol.MemoryUser, protocol.Timestamp, bool, error)
	VMessTimedUserValidatorGetAEAD(userHash []byte) (*protocol.MemoryUser, bool, error)
	VMessGetUser(email string) *protocol.MemoryUser
	ShadowsocksValidatorGet(bs []byte, command protocol.RequestCommand) (mu *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error)
}

var authenticators = make([]Authenticator, 4)

func RegisterAuthenticator(a Authenticator) {
	authenticators = append(authenticators, a)
}

func ExtraAuthenticationUsed() bool {
	return len(authenticators) > 0
}
