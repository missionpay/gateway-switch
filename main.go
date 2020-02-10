package main

import (
	"fmt"
	"os"

	"github.com/kyani-inc/kms-ecs-template/api"
	log "github.com/kyani-inc/kms-ecs-template/logger"
	"github.com/labstack/echo/v4"
)

const logo = `    __ __      _  _        _ 
   / //_/_  __(_)(_)____  (_)
  / ,< / / / / __ / __  \/ /
 / /| / /_/ / /_/ / / / / /
/_/ |_\__, /\__,_/_/ /_/_/
      __/ / Live Kyani
    /____/
`

// Bootstrap the App
func init() {
}

func main() {
	// Echo instance
	e := echo.New()

	// Health Info
	health := map[string]string{
		"app":   os.Getenv("APP"),
		"build": os.Getenv("BUILD"),
	}

	api.Routes(e)

	// Health endpoint
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(200, health)
	})

	// Start server
	e.HideBanner = true
	fmt.Println(logo)
	err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// If we get to this point Log any errors and shut it down
	if err != nil {
		log.Fatal("Failed to start Server", err.Error())
	}
}
