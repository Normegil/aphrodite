package security

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/normegil/aphrodite/model"
	"github.com/normegil/aphrodite/modules/datasource"
)

const AUTHORIZATION_SEPARATOR = ":"

type Authenticator struct {
	datasource.SpecificUserLoader
}

func (a *Authenticator) Authenticate(r *http.Request) (*model.User, error) {
	// Check for JWT
	return a.headersAuthentication(r)
}

func (a *Authenticator) headersAuthentication(r *http.Request) (*model.User, error) {
	authType := r.Header.Get("WWW-Authenticate")
	if "" == authType {
		return nil, nil
	}
	if "Basic" == authType {
		return a.basicAuth(r.Header.Get("Authorization"))

	} else if "Digest" == authType {
		return nil, fmt.Errorf("Authentication protocol not supported (Got from WWW-Authenticate header): %s", authType)
	}
	return nil, fmt.Errorf("Authentication protocol not recognized (Got from WWW-Authenticate header): %s", authType)
}

var AuthenticationFailed = errors.New("Authentication Failed")

func (a *Authenticator) basicAuth(authorizationHeader string) (*model.User, error) {
	decoded, err := base64.StdEncoding.DecodeString(authorizationHeader)
	if nil != err {
		return nil, err
	}
	header := strings.SplitN(string(decoded), AUTHORIZATION_SEPARATOR, 2)
	name := header[0]
	pass := header[1]

	user, err := a.User(name)
	if nil != err {
		return nil, err
	}

	success, err := user.Password().Check(pass)
	if nil != err {
		return nil, err
	}
	if success {
		return user, nil
	}
	return nil, AuthenticationFailed
}
