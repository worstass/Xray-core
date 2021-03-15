package auth

import (
	"crypto/cipher"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/uuid"
)

type Authenticator interface {
	VMessGetUser(email string) *protocol.MemoryUser
	VLessGet(id uuid.UUID) *protocol.MemoryUser
	TrojanGet(hash string)  *protocol.MemoryUser
	ShadowsocksCount() int
	ShadowsocksGet(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error)
}

type NullAuthenticator struct {}

func (a *NullAuthenticator) ShadowsocksCount() int {
	return 2
}

func (a *NullAuthenticator) ShadowsocksGet(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	return nil, nil, nil, 0, err
}

func (a *NullAuthenticator) SSCount() int {
	return 2
}

func (a *NullAuthenticator) VMessGetUser(email string) *protocol.MemoryUser {
	return nil
}

func (a *NullAuthenticator) VLessGet(id uuid.UUID) *protocol.MemoryUser {
	return nil
}

func (a *NullAuthenticator) TrojanGet(hash string)  *protocol.MemoryUser {
	return nil
}

var authenticator Authenticator = &NullAuthenticator{}

func SetAuthenticator(a Authenticator)  {
	authenticator = a
}

func ShouldNotBeCalled() {
	panic("Should not be called")
}