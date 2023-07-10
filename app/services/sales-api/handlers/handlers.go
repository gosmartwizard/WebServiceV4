package handlers

import (
	"net/http"
	"os"

	"github.com/gosmartwizard/WebServiceV4/app/services/sales-api/handlers/v1/testgrp"
	"github.com/gosmartwizard/WebServiceV4/business/web/v1/mid"
	"github.com/gosmartwizard/WebServiceV4/foundation/web"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log))

	app.Handle(http.MethodGet, "/test", testgrp.Test)

	return app
}
