package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/aldiramdan/go-backend/libs"
)

func AuthMidle(role ...string) Middle {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			var header string
			var isAllowed bool

			if header = r.Header.Get("Authorization"); header == "" {
				libs.GetResponse("header not provide, please login", 401, true).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				libs.GetResponse("invalid header type", 401, true).Send(w)
				return
			}

			token := strings.Replace(header, "Bearer ", "", -1)

			checkToken, err := libs.VerifyToken(token)
			if err != nil {
				libs.GetResponse(err.Error(), 401, true).Send(w)
				return
			}

			for _, role := range role {
				if role == checkToken.Role {
					isAllowed = true
				}
			}

			if !isAllowed {
				libs.GetResponse("you not have permission to accsess", 401, true).Send(w)
				return
			}

			ctx := context.WithValue(r.Context(), "user", checkToken.UserId)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
