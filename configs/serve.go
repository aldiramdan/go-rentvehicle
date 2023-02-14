package configs

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aldiramdan/go-backend/routers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start application",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	t := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return t
}

func serve(cmd *cobra.Command, args []string) error {

	mainRoute, err := routers.IndexRoute()
	if err != nil {
		return err
	}

	var address string = "0.0.0.0:8080"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "0.0.0.0:" + PORT
	}

	corss := corsHandler()

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
		Handler:      corss.Handler(mainRoute),
	}

	log.Println("app run on port", address)

	return srv.ListenAndServe()

}
