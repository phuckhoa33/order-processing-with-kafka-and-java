package main

import (
	application "github.com/phuckhoa33/web-crawler"
	"github.com/phuckhoa33/web-crawler/internal/config"
)

// @version		1.0
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// Init config
	config := config.NewConfig()

	// Start app
	application.Start(config)
}
