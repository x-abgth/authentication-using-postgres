package utils

import (
	"github.com/gorilla/sessions"
)

var (
	Key   string
	Store = sessions.NewCookieStore([]byte(Key))
)
