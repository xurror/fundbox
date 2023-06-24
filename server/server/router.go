package server

import (
	"flag"
	"getting-to-go/controller"
	"getting-to-go/model"
	fmiddleware "getting-to-go/server/middleware"
	"getting-to-go/service"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	jwtauth "github.com/go-chi/jwtauth/v5"
	"time"
)

func NewRouter() chi.Router {
	flag.Parse()
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(CorsOptions()))

	userService := service.NewUserService()
	authService := service.NewAuthService(userService)
	authController := controller.NewAuthController(authService)

	r.Route("/auth", authController.Register())
	r.Handle("/graphiql", graphiQlHandler("Fund Box", "/query"))
	r.Group(func(r chi.Router) {
		r.Use(fmiddleware.Verifier(jwtauth.New("HS256", []byte("secret"), nil), model.DB()))
		r.Use(jwtauth.Authenticator)
		r.Handle("/query", graphqlHandler())
	})

	return r
}
