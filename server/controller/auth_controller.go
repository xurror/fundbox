package controller

import (
	"errors"
	"getting-to-go/service"
	_type "getting-to-go/type"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// AuthController provides user-related endpoints
type AuthController struct {
	authService *service.AuthService
}

// NewAuthController creates a new AuthController instance
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", c.login)
		r.Post("/signup", c.signUp)
	}
}

func (c *AuthController) login(w http.ResponseWriter, r *http.Request) {
	login := &LoginRequest{}
	if err := render.Bind(r, login); err != nil {
		render.Render(w, r, _type.ErrInvalidRequest(err))
		return
	}

	user, err := c.authService.Authenticate(login.Email, login.Password)
	if err != nil {
		render.Render(w, r, _type.ErrUnauthorized(err))
		return
	}

	now := time.Now()
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	token, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     now.Add(time.Minute * 15).Unix(),
		"iat":     now.Unix(),
	})

	render.Status(r, http.StatusOK)
	render.Render(w, r, &JwtResponse{
		Code:   http.StatusOK,
		Expire: token.Expiration(),
		Token:  tokenString,
	})
}

func (c *AuthController) signUp(w http.ResponseWriter, r *http.Request) {
	req := &SignUpRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, _type.ErrInvalidRequest(err))
		return
	}

	user, err := c.authService.SignUp(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		req.Roles,
	)
	if err != nil {
		switch e := err.(type) {
		case _type.ErrResponse:
			render.Render(w, r, _type.AppErrRender(e))
		default:
			render.Render(w, r, _type.ErrInternal(err))
		}
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Roles:     user.Roles,
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginRequest) Bind(r *http.Request) error {
	if l.Email == "" {
		return errors.New("email is required")
	}
	if l.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type JwtResponse struct {
	Code   int       `json:"code"`
	Expire time.Time `json:"expire"`
	Token  string    `json:"token"`
}

func (j *JwtResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type SignUpRequest struct {
	FirstName string   `json:"first_name" binding:"required"`
	LastName  string   `json:"last_name" binding:"required"`
	Email     string   `json:"email" binding:"required,email"`
	Password  string   `json:"password" binding:"required,min=6"`
	Roles     []string `json:"roles" binding:"required"`
}

func (s *SignUpRequest) Bind(r *http.Request) error {
	if s.FirstName == "" {
		return errors.New("first_name is required")
	}
	if s.LastName == "" {
		return errors.New("last_name is required")
	}
	if s.Email == "" {
		return errors.New("email is required")
	}
	if s.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Roles     []string  `json:"roles"`
}

func (j *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
