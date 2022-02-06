package httputil

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func CustomerTokenMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return Handler(func(w http.ResponseWriter, r *http.Request) error {
			userID := r.Header.Get(HeaderUserID)
			if userID == "" {
				return errors.New("auth: expected user id")
			}

			tokenID := r.Header.Get(HeaderTokenID)
			if tokenID == "" {
				return errors.New("auth: expected token id")
			}

			tokenName := r.Header.Get(HeaderTokenName)
			if tokenName == "" {
				return errors.New("auth: expected token name")
			}

			ctx := WithUserID(r.Context(), userID)
			ctx = WithTokenID(ctx, tokenID)
			ctx = WithTokenName(ctx, tokenName)

			next.ServeHTTP(w, r.WithContext(ctx))

			return nil
		})
	}
}
