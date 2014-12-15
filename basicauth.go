package app

import "net/http"

func BasicAuth(username, password string) Middleware {
	return func(inner http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := http.BasicAuth(r)
			if !ok || user != username || pass != password {
				setBasicAuth(w)
				return
			}
			inner.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
