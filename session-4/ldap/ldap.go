package ldap

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/go-ldap/ldap"
)

const (
	ldapServer   = "ldap.forumsys.com"
	ldapPort     = 389
	ldapBindDN   = "cn=read-only-admin,dc=example,dc=com"
	ldapPassword = "password"
	ldapSearchDN = "dc=example,dc=com"
)

type UserLDAPData struct {
	ID       string
	Email    string
	Name     string
	FullName string
}

func AuthUsingLDAP(username, password string) (bool, *UserLDAPData, error) {

	// init ldap connection
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		log.Error(err)
		return false, nil, err
	}
	defer l.Close()

	// bind to ldap server
	if err = l.Bind(ldapBindDN, ldapPassword); err != nil {
		return false, nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		ldapSearchDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn", "cn", "sn", "mail"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	if len(sr.Entries) == 0 {
		return false, nil, errors.New("user not found")
	}

	entry := sr.Entries[0]
	if err = l.Bind(entry.DN, password); err != nil {
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
