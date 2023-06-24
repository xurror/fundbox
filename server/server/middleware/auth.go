package middleware

import (
	"context"
	"getting-to-go/model"
	_type "getting-to-go/type"
	jwtauth "github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Verifier(ja *jwtauth.JWTAuth, db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := jwtauth.TokenFromHeader(r)
			token, err := jwtauth.VerifyToken(ja, tokenString)
			ctx := jwtauth.NewContext(r.Context(), token, err)

			if err != nil {
				render.Render(w, r, _type.ErrInvalidRequest(err))
				return
			}

			email, _ := token.Get("email")
			user := &model.User{}
			result := db.First(&user, "email = ?", email)
			if result.Error != nil {
				render.Render(w, r, _type.ErrInvalidRequest(err))
				return
			}

			// Add the user to the context
			ctx = context.WithValue(ctx, "userId", user.ID)
			ctx = context.WithValue(ctx, "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
