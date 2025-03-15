package auth

import "net/http"

type AuthHandlerDep struct {
	AuthService *AuthService
}

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, dep AuthHandlerDep) {
	handler := &AuthHandler{
		AuthService: dep.AuthService,
	}

	router.HandleFunc("/login", handler.Login())
	router.HandleFunc("/register", handler.Register())
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login logics is not created yet\n"))
		w.Write([]byte(a.AuthService.SayHello()))
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Registration logics is not created yet\n"))
		w.Write([]byte(a.AuthService.SayHello()))
	}
}
