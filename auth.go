package main

import (
	"net/http"
)

type Auth struct {
	Storage *Storage
}

func (auth *Auth) checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, ok := auth.Storage.GetUser(username)

		if !ok || user.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
