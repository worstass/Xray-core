package auth

import (
	"github.com/xtls/xray-core/common/protocol"
)

func VMessGetUser(email string) *protocol.MemoryUser {
	for _, a := range authenticators {
		res := a.VMessGetUser(email)
		if res != nil {
			return res
		}
	}
	return nil
}

func VMessTimedUserValidatorGet(userHash []byte) (mu *protocol.MemoryUser, ts protocol.Timestamp, b bool, err error) {
	for _, a := range authenticators {
		mu, ts, b, err := a.VMessTimedUserValidatorGet(userHash)
		if mu != nil {
			return mu, ts, b, err
		}
	}
	return nil, ts, b, err
}

func VMessTimedUserValidatorGetAEAD(userHash []byte) (mu *protocol.MemoryUser, b bool, err error) {
	for _, a := range authenticators {
		mu, b, err := a.VMessTimedUserValidatorGetAEAD(userHash)
		if mu != nil {
			return mu, b, err
		}
	}
	return nil, b, err
}