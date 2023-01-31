package main

import (
	"log"
	"net/http"

	"github.com/aldiramdan/go-backend/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	r, err := routers.IndexRoute()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("app run on port :3080")
	err = http.ListenAndServe("127.0.0.1:3080", r)

	if err != nil {
		log.Fatal(err)
	}
	
}
