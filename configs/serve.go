package configs

import (
	"log"
	"net/http"
	"os"

	"github.com/aldiramdan/go-backend/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start application",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {

	mainRoute, err := routers.IndexRoute()
	if err != nil {
		return err
	}

	var address string = "127.0.0.1:8080"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "127.0.0.1:" + PORT
	}

	log.Println("app run on port", address)

	return http.ListenAndServe(address, mainRoute)

}
