package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/database"
	"github.com/Muh-Sidik/kasir-api/docs"
	_ "github.com/Muh-Sidik/kasir-api/docs"
	"github.com/Muh-Sidik/kasir-api/internal/route"
)

// @title Swagger Kasir API
// @version 1.0
// @description This is a kasir server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	e := config.LoadConfig()

	docs.SwaggerInfo.Host = e.APP_HOST + ":" + e.APP_PORT
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	db := database.New(e)
	defer db.Close()

	mux := http.NewServeMux()

	route.Setup(mux, e, db)

	fmt.Println("Successfully listen server in port :8000")
	err := http.ListenAndServe(
		":8000",
		mux,
	)

	if err != nil {
		log.Fatalf("error server: %v", err)
	}

}
