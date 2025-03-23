package main

import (
	"log"
	"matchfinder"
	"matchfinder/internal/auth"
	"matchfinder/internal/user"
	"matchfinder/pkg/db"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

const (
	confPath = "C:\\Son_Alex\\GO_projects\\MatchFinder\\MatchFinder\\configs"
)

func main() {
	//----------------------------DB--------------------------------
	db, err := db.NewPostgresDB(db.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "my_pass",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("Failed to initialize db: %s\n", err.Error())
	}

	router := http.NewServeMux()
	if err := initConfig(); err != nil {
		log.Fatalf("An error during config reading: %s\n", err.Error())
	}
	//----------------------------DB---------------------------------

	//------------------------REPOSITORY-----------------------------
	userRepository := user.NewUserRepository(db)
	//------------------------REPOSITORY-----------------------------

	//-------------------------SERVICES------------------------------
	authService := auth.NewAuthService(userRepository)
	//-------------------------SERVICES------------------------------

	//-------------------------HANDLERS------------------------------
	auth.NewAuthHandler(router, auth.AuthHandlerDep{
		AuthService: authService,
	})
	//-------------------------HANDLERS------------------------------

	//--------------------------SERVER-------------------------------
	srv := matchfinder.NewServer(viper.GetString("port"), router)
	err = srv.Run()
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
