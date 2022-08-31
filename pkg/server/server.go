package pkg

import (
	"log"
	"net/http"

	"github.com/n4ze3m/termshare/pkg/server/routes"
)

func ServerInit() {
	log.Println("Starting server...")
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
