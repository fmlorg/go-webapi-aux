package http

import (
	"net/http"
	"strconv"
	"time"
)

var defaultSessionKey string = "_session"

func SetSessionCookie(w http.ResponseWriter, r *http.Request, key string, val string) {
	// the 5th arg "expiryDays == 0" implies the session cookie
	SetCookie(w, r, key, val, 0)
}

func SetCookie(w http.ResponseWriter, r *http.Request, key string, val string, expiryDays int) {
	if len(key) == 0 {
		key = defaultSessionKey
	}
	if len(val) == 0 {
		now := time.Now().UnixNano()
		val = strconv.FormatInt(now, 10)
	}

	now := time.Now()
	expired := now.AddDate(0, 0, expiryDays)
	cookie := &http.Cookie{
		Name:    key,
		Value:   val,
		Expires: expired,
	}

	// the session cookie if expiryDays == 0
	// too old Expires date (year < 1601) is invalid.
	// http.cookie does not add Expires: header field if it is set.
	// See http.cookie source code for more details.
	if expiryDays == 0 {
		cookie.Expires = now.AddDate(-1000, 0, 0)
	}

	http.SetCookie(w, cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request, key string) string {
	if len(key) == 0 {
		key = defaultSessionKey
	}

	token, err := r.Cookie(key)
	if err != nil {
		return ""
	}
	return token.Value
}
