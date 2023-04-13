package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should sessions last?
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookies persist?
	persist = false
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}

	// should cookies be secure
	secure = true
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	switch strings.ToLower(c.SessionType) {
	case "redis":
		// TODO
	case "mysql", "mariadb":
		// TODO
	case "postgres", "postgresql":
		// TODO
	default:
		// TODO
	}

	return session
}
