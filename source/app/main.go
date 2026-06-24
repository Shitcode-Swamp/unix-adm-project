package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	appctx "github.com/Shitcode-Swamp/unix-adm-project/source/app/context"
	"github.com/Shitcode-Swamp/unix-adm-project/source/controller"
)

func main() {
	app, err := appctx.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Seed(context.Background()); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	controller.Setup(r, app)

	addr := ":" + appctx.Cfg.ServerPort

	if appctx.Cfg.TLSCert != "" && appctx.Cfg.TLSKey != "" {
		log.Fatal(r.RunTLS(addr, appctx.Cfg.TLSCert, appctx.Cfg.TLSKey))
	} else {
		log.Fatal(r.Run(addr))
	}
}
