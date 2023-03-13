package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/ikhwankhaliddd/shorten-url-gofiber-redis/database"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rIncremental := database.CreateClient(1)
	defer rIncremental.Close()

	_ = rIncremental.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}
