package auth

import (
	"crypto/cipher"
	"github.com/xtls/xray-core/common/protocol"
)

func ShadowsocksValidatorGet(bs []byte, command protocol.RequestCommand) (mu *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	for _, a := range authenticators {
		mu, aead, ret, ivLen, err := a.ShadowsocksValidatorGet(bs, command)
		if mu != nil {
			return mu, aead, ret, ivLen, err
		}
	}
	return nil, aead, ret, ivLen, err
}