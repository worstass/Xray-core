package auth

import (
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/uuid"
)

type Authenticator interface {
	VLessGet(id uuid.UUID) *protocol.MemoryUser
	TrojanGet(hash string)  *protocol.MemoryUser
    VMessTimedUserValidatorGet(userHash []byte) (*protocol.MemoryUser, protocol.Timestamp, bool, error)
    VMessTimedUserValidatorGetAEAD(userHash []byte) (*protocol.MemoryUser, bool, error)
	VMessGetUser(email string) *protocol.MemoryUser
}

type NullAuthenticator struct {}

func (a *NullAuthenticator) VMessTimedUserValidatorGet(userHash []byte) (*protocol.MemoryUser, protocol.Timestamp, bool, error) {
	return nil, 0, false, nil
}

func (a *NullAuthenticator) VMessTimedUserValidatorGetAEAD(userHash []byte) (*protocol.MemoryUser, bool, error) {
	return nil, false, nil
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

type errPathObjHolder struct{}

func ShouldNotBeCalled() error {
	err := errors.New("Should not be called")
	err.WithPathObj(errPathObjHolder{}).AtDebug().WriteToLog()
	return err
}