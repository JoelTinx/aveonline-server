package main

import (
	cmd "github.com/JoelTinx/aveonline-server/cmd/rest"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/JoelTinx/aveonline-server/docs"
)

// @title API Facturacion
// @version 1.0.0
// @description Permite gestionar la facturación de Aveonline.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))
	app.Get("/swagger/*", swagger.HandlerDefault)
	cmd.SetupRoutes(app)
	app.Listen(":5000")
}
