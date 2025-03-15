package main

import (
	"log"
	"matchfinder"
	"matchfinder/internal/auth"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	//-------------------------SERVICES------------------------------
	authService := auth.NewAuthService()
	//-------------------------SERVICES------------------------------

	//-------------------------HANDLERS------------------------------
	auth.NewAuthHandler(router, auth.AuthHandlerDep{
		AuthService: authService,
	})
	//-------------------------HANDLERS------------------------------

	//--------------------------SERVER-------------------------------
	srv := matchfinder.NewServer("8080", router)
	err := srv.Run()
	if err != nil {
		log.Printf("Some error during server starting: %s", err.Error())
	}
	//--------------------------SERVER-------------------------------
}
