package main

import (
	"log"
	"matchfinder"
	"matchfinder/internal/auth"
	"net/http"

	"github.com/spf13/viper"
)

const (
	confPath = "C:\\Son_Alex\\GO_projects\\MatchFinder\\MatchFinder\\configs"
)

func main() {
	router := http.NewServeMux()
	if err := initConfig(); err != nil {
		log.Fatalf("An error during config reading: %s\n", err.Error())
	}

	//-------------------------SERVICES------------------------------
	authService := auth.NewAuthService()
	//-------------------------SERVICES------------------------------

	//-------------------------HANDLERS------------------------------
	auth.NewAuthHandler(router, auth.AuthHandlerDep{
		AuthService: authService,
	})
	//-------------------------HANDLERS------------------------------

	//--------------------------SERVER-------------------------------
	srv := matchfinder.NewServer(viper.GetString("port"), router)
	err := srv.Run()
	if err != nil {
		log.Printf("Some error during server starting: %s", err.Error())
	}
	//--------------------------SERVER-------------------------------
}

func initConfig() error {
	viper.AddConfigPath(confPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
