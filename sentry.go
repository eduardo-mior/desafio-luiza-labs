package main

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func initSentry() {
	sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		AttachStacktrace: true,
	})
	defer sentry.Flush(2 * time.Second)
}

func setupSentryMiddleware(router *gin.Engine) {
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
}
