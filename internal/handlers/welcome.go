package handlers   // this will say that welcome.go is part of handler package

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error{   // Ctx is context and struct
    return c.Render("welcome",nil,"layouts/main")
}

