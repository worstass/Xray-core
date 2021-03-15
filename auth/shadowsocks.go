package auth

import (
	"crypto/cipher"
	"github.com/xtls/xray-core/common/protocol"
)

func ShadowsocksCount() int  {
	return authenticator.ShadowsocksCount()
}

func ShadowsocksGet(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	return authenticator.ShadowsocksGet(bs, command)
}