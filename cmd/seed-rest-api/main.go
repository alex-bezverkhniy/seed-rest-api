package main

import (
	_ "seed-rest-api/docs"
	"seed-rest-api/internal/infrastructure"
)

// @title API
// @version 1.0
// @description This is Seed REST API Docs.
// @contact.name Alex Bezverkhniy
// @contact.email alexandr.bezverkhniy@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	infrastructure.Run()
}
