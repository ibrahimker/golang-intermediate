package service

import (
	"errors"

	"github.com/ibrahimker/golang-intermediate/session-4/repository"
	log "github.com/sirupsen/logrus"
)

type LoginSvc struct {
	loginRepo *repository.LDAPRepo
}

func NewLoginService(loginRepo *repository.LDAPRepo) *LoginSvc {
	return &LoginSvc{loginRepo: loginRepo}
}

func (s *LoginSvc) Authenticate(username, password string) (*repository.UserLDAPData, error) {
	ok, data, err := s.loginRepo.AuthUsingLDAP(username, password)
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
