package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	appctx "github.com/Shitcode-Swamp/unix-adm-project/source/app/context"
	"github.com/Shitcode-Swamp/unix-adm-project/source/controller"
)

func main() {
	app, err := appctx.New()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	controller.Setup(r, app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
