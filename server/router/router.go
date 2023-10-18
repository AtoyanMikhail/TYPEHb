package router

import (
	"math/rand"
	"typenb/textgen"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app           *fiber.App
	textGenerator textgen.Generator
}

func (r Router) initHandlers() {
	//

	r.app.Get("/text/", func(c *fiber.Ctx) error {
		seed := rand.Int63()
		text := r.textGenerator.Generate(seed)

		return c.SendString(text) // r.textGenerator
	})

	r.app.Get("/text/:seed", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") // тут seed получишь с Params
	})

}

func (r Router) Start(addr string) {
	r.app.Listen(addr)
}

func New(textGenerator textgen.Generator) Router {
	router := Router{
		app:           fiber.New(),
		textGenerator: textGenerator,
	}

	router.initHandlers()

	return router
}
