package shadowsocks

import (
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"hash/crc64"
	"strings"
	"sync"

<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/xtls/xray-core/app/extra/auth"
=======
	"github.com/xtls/xray-core/common/dice"
>>>>>>> c3298c38a0d6f9c66703a6dd565e783778de8b35
=======
>>>>>>> c4a3dbdeac05220d8d06e07467f9cf9fd356fd60
	"github.com/xtls/xray-core/common/protocol"
)

// Validator stores valid Shadowsocks users.
type Validator struct {
	sync.RWMutex
	users []*protocol.MemoryUser

	behaviorSeed  uint64
	behaviorFused bool
}

var (
	ErrNotFound = newError("Not Found")
)

// Add a Shadowsocks user.
func (v *Validator) Add(u *protocol.MemoryUser) error {
<<<<<<< HEAD
<<<<<<< HEAD
	if !auth.NoAuthenticator {
		return auth.ShouldNotBeCalled()
	}

=======
>>>>>>> c4a3dbdeac05220d8d06e07467f9cf9fd356fd60
	account := u.Account.(*MemoryAccount)
=======
	v.Lock()
	defer v.Unlock()
>>>>>>> c3298c38a0d6f9c66703a6dd565e783778de8b35

	account := u.Account.(*MemoryAccount)
	if !account.Cipher.IsAEAD() && len(v.users) > 0 {
		return newError("The cipher is not support Single-port Multi-user")
	}
	v.users = append(v.users, u)

	if !v.behaviorFused {
		hashkdf := hmac.New(sha256.New, []byte("SSBSKDF"))
		hashkdf.Write(account.Key)
		v.behaviorSeed = crc64.Update(v.behaviorSeed, crc64.MakeTable(crc64.ECMA), hashkdf.Sum(nil))
	}

	return nil
}

// Del a Shadowsocks user with a non-empty Email.
<<<<<<< HEAD
func (v *Validator) Del(e string) error {
	if e == "" {
=======
func (v *Validator) Del(email string) error {
	if email == "" {
>>>>>>> c3298c38a0d6f9c66703a6dd565e783778de8b35
		return newError("Email must not be empty.")
	}

<<<<<<< HEAD
// Count the number of Shadowsocks users
func (v *Validator) Count() int {
	length := 0
	v.users.Range(func(_, _ interface{}) bool {
		length++

		return true
	})
	return length
}

// Get a Shadowsocks user and the user's cipher.
func (v *Validator) Get(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	var dataSize int
=======
	v.Lock()
	defer v.Unlock()

	email = strings.ToLower(email)
	idx := -1
	for i, u := range v.users {
		if strings.EqualFold(u.Email, email) {
			idx = i
			break
		}
	}
>>>>>>> c3298c38a0d6f9c66703a6dd565e783778de8b35

	if idx == -1 {
		return newError("User ", email, " not found.")
	}
	ulen := len(v.users)

	v.users[idx] = v.users[ulen-1]
	v.users[ulen-1] = nil
	v.users = v.users[:ulen-1]

	return nil
}

// Get a Shadowsocks user.
func (v *Validator) Get(bs []byte, command protocol.RequestCommand) (u *protocol.MemoryUser, aead cipher.AEAD, ret []byte, ivLen int32, err error) {
	v.RLock()
	defer v.RUnlock()

	for _, user := range v.users {
		if account := user.Account.(*MemoryAccount); account.Cipher.IsAEAD() {
			aeadCipher := account.Cipher.(*AEADCipher)
			ivLen = aeadCipher.IVSize()
			iv := bs[:ivLen]
			subkey := make([]byte, 32)
			subkey = subkey[:aeadCipher.KeyBytes]
			hkdfSHA1(account.Key, iv, subkey)
			aead = aeadCipher.AEADAuthCreator(subkey)

			var matchErr error
			switch command {
			case protocol.RequestCommandTCP:
				data := make([]byte, 16)
				ret, matchErr = aead.Open(data[:0], data[4:16], bs[ivLen:ivLen+18], nil)
			case protocol.RequestCommandUDP:
				data := make([]byte, 8192)
				ret, matchErr = aead.Open(data[:0], data[8180:8192], bs[ivLen:], nil)
			}

			if matchErr == nil {
				u = user
				err = account.CheckIV(iv)
				return
			}
		} else {
			u = user
			ivLen = user.Account.(*MemoryAccount).Cipher.IVSize()
			// err = user.Account.(*MemoryAccount).CheckIV(bs[:ivLen]) // The IV size of None Cipher is 0.
			return
		}
	}

	return nil, nil, nil, 0, ErrNotFound
}

<<<<<<< HEAD
// Get the only user without authentication
func (v *Validator) GetOnlyUser() (u *protocol.MemoryUser, ivLen int32) {
	v.users.Range(func(_, user interface{}) bool {
		u = user.(*protocol.MemoryUser)
		return false
	})
	ivLen = u.Account.(*MemoryAccount).Cipher.IVSize()
<<<<<<< HEAD
=======
func (v *Validator) GetBehaviorSeed() uint64 {
	v.Lock()
	defer v.Unlock()
>>>>>>> c3298c38a0d6f9c66703a6dd565e783778de8b35

	v.behaviorFused = true
	if v.behaviorSeed == 0 {
		v.behaviorSeed = dice.RollUint64()
	}
	return v.behaviorSeed
}
=======

	return
}
>>>>>>> c4a3dbdeac05220d8d06e07467f9cf9fd356fd60
