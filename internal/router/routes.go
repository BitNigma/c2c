package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store *session.Store
)

func Setup(app *fiber.App) {
	/* Sessions Config */
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration:   time.Hour * 2,
		CookieSecure: true,
	})

	// Static files
	app.Static("/static", "./static")

	/* Views */
	app.Get("/", CheckAuth(HandleHome))
	app.Get("/login", HandleLogin)
	app.Post("/login", HandleLoginAuth)
	app.Get("/signup", HandleSignUp)
	app.Post("/signup", HandleRegister)
	app.Get("/recover", HandleRecover)
	app.Post("/recover", HandleRecover)
	app.Use("/logout", HandleLogout, HandleHome)

	/* Views protected with session middleware */
	personal := app.Group("/personal", AuthMiddleware)
	personal.Get("/kyc", HandleKyc)
	personal.Get("/payments", HandlePayments)

	/* Page Not Found Management */

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./static/page-error-404.html")
	})
}
