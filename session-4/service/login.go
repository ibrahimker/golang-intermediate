package service

import (
	"errors"

	"github.com/ibrahimker/golang-intermediate/session-4/ldap"
	log "github.com/sirupsen/logrus"
)

func Authenticate(username, password string) (*ldap.UserLDAPData, error) {
	ok, data, err := ldap.AuthUsingLDAP(username, password)
	if !ok {
		err := errors.New("auth using ldap not ok")
		log.Error("auth using ldap not ok")
		return nil, err
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return data, nil
}
