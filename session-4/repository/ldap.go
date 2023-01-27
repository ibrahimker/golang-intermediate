package repository

import (
	"errors"
	"fmt"

	"github.com/ibrahimker/golang-intermediate/session-4/config"
	log "github.com/sirupsen/logrus"

	"github.com/go-ldap/ldap"
)

type UserLDAPData struct {
	ID       string
	Email    string
	Name     string
	FullName string
}

type LDAPRepo struct {
	conn *ldap.Conn
}

func NewLDAPRepo(conn *ldap.Conn) *LDAPRepo {
	return &LDAPRepo{conn: conn}
}

func (r *LDAPRepo) AuthUsingLDAP(username, password string) (bool, *UserLDAPData, error) {
	searchRequest := ldap.NewSearchRequest(
		config.LdapSearchDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn", "cn", "sn", "mail"},
		nil,
	)

	sr, err := r.conn.Search(searchRequest)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	if len(sr.Entries) == 0 {
		return false, nil, errors.New("user not found")
	}

	entry := sr.Entries[0]
	if err = r.conn.Bind(entry.DN, password); err != nil {
		log.Error(err)
		return false, nil, err
	}

	data := &UserLDAPData{ID: username}
	for _, attr := range entry.Attributes {
		switch attr.Name {
		case "sn":
			data.Name = attr.Values[0]
		case "mail":
			data.Email = attr.Values[0]
		case "cn":
			data.FullName = attr.Values[0]
		}
	}

	return true, data, nil
}
