package auth

import (
	"crypto/cipher"
	"github.com/xtls/xray-core/common/protocol"
)

func VMessGetUser(email string) *protocol.MemoryUser {
	return authenticator.VMessGetUser(email)
}

func VMessTimedUserValidatorGet(userHash []byte) (*protocol.MemoryUser, protocol.Timestamp, bool, error) {
	return authenticator.VMessTimedUserValidatorGet(userHash)
}

func VMessTimedUserValidatorGetAEAD(userHash []byte) (*protocol.MemoryUser, bool, error) {
	return authenticator.VMessTimedUserValidatorGetAEAD(userHash)
}

func ShadowsocksValidatorGet(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	return authenticator.ShadowsocksValidatorGet(bs, command)
}