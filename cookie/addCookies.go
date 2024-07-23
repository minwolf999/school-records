package cookie

import (
	"net/http"
)

// Create and set a session cookie for the user
func AddCookies(w http.ResponseWriter, value string) {
	cookieEmail := http.Cookie{
		Name:     "cookie",
		Value:    value,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookieEmail)
}
