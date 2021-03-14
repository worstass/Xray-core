package auth

type Authenticator interface {
	 SSCount() int
}

type NullAuthenticator struct {}

func (a *NullAuthenticator) SSCount() int {
	return 2
}

var authenticator Authenticator = &NullAuthenticator{}

func SetAuthenticator(a Authenticator)  {
	authenticator = a
}

func ShouldNotBeCalled() {
	panic("Should not be called")
}