package main

import (
	"fmt"
	"sagara-try/config"
	"sagara-try/helpers"
	"sagara-try/middleware"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helpers.Logger("error", "Error getting env")
	}

	config.Init()

	ctm := mux.NewRouter()
	ctm.Use(middleware.JwtAuth)

	presenter := Factory(ctm)
	fmt.Println(presenter)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	p := os.Getenv("PORT")
	h := c.Handler(ctm)
	s := new(http.Server)
	s.Handler = h
	s.Addr = ":" + p
	appASCIIArt := fmt.Sprintf(`
		__________ __________ _________ 
		/   /_____//   /_____//    O___ \
		\___\%%%%%%\___\%%%%%%\_____\  \/
		 "EEEEEEE" "EEEEEEEE" "EEEEE"\__\

		Developed by Lutfi M 
		Server run in port %s
	`, s.Addr)
	fmt.Println(appASCIIArt)
	s.ListenAndServe()
}
