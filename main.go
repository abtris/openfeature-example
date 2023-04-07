package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature"
	"github.com/rs/zerolog"
)

const defaultMessage = "Hello!"
const newWelcomeMessage = "Hello, welcome to this OpenFeature-enabled website!"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zerologr.SetMaxV(1)

	zl := zerolog.New(os.Stderr)
	zl = zl.With().Caller().Timestamp().Logger()
	var log logr.Logger = zerologr.New(&zl)

	openfeature.SetLogger(log)
	openfeature.SetProvider(flagd.NewProvider(flagd.WithoutCache()))
	client := openfeature.NewClient("GoStartApp").WithLogger(log)

	// Initialize Go Gin
	engine := gin.Default()

	// Setup a simple endpoint
	engine.GET("/hello", func(c *gin.Context) {
		welcomeMessage, _ := client.BooleanValue(
			context.Background(), "welcome-message", false, openfeature.EvaluationContext{},
		)
		if welcomeMessage {
			c.JSON(http.StatusOK, newWelcomeMessage)
			return
		} else {
			c.JSON(http.StatusOK, defaultMessage)
			return
		}
	})

	engine.Run()
}
