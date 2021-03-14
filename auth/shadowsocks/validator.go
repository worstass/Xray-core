package shadowsocks

import (
	"crypto/cipher"
	"github.com/xtls/xray-core/common/protocol"
)

func Count() int  {
	return 2
}

func Get(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	return nil, nil, nil, 0, err
}