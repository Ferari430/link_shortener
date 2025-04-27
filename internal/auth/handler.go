package auth

import (
	"fmt"
	"log"
	"my_project/configs"
	"my_project/pkg/jwt"
	"my_project/pkg/req"
	"my_project/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	Config *configs.Config
	Auth   *AuthService
}

type AuthHandler struct {
	Config *configs.Config
	Auth   *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
		Auth:   deps.Auth,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r) // ИЗБАВИТЬСЯ ОТ ДЖЕНЕРИКОВ И ПЕРЕПИСАТЬ ПО ЧЕЛОВЕЧЕСКИ
		if err != nil {
			return
		}

		email, err := handler.Auth.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		data := LoginResponce{
			Token: token,
		}
		_ = data
		res.Json(w, data, 200)
		log.Println(email)

		fmt.Printf("email is %s and password is %s\n", body.Email, body.Password)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register Handler")

		body, err := req.HandleBody[RegisterRequest](&w, r) // ИЗБАВИТЬСЯ ОТ ДЖЕНЕРИКОВ И ПЕРЕПИСАТЬ ПО ЧЕЛОВЕЧЕСКИ
		if err != nil {
			return
		}

		email, err := handler.Auth.Register(body.Email, body.Password, body.Name)

		if err != nil {
			http.Error(w, "Cant Register", 401)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		data := RegisterResponce{Token: token}
		log.Println(data)
		res.Json(w, data, 200)
		fmt.Printf("Registration succsess - email is %s and password is %s, Name is %s\n", email, body.Password, body.Name)

	}
}
