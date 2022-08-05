package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"go-study/go-ent-pokemon/ent"
	"log"
)

func main() {

	viper.SetConfigFile("go-ent-pokemon/.env")
	viper.ReadInConfig()

	ctx := context.Background()

	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True",
		viper.Get("DB_USER"),
		viper.Get("DB_PASSWORD"),
		viper.Get("DB_HOST"),
		viper.Get("DB_PORT"),
		viper.Get("DB_NAME"))

	// connect to mysql
	client, err := ent.Open(dialect.MySQL, url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/create/battle", func(c *fiber.Ctx) error {
		payload := struct {
			Result    string `json:"result"`
			Contender int    `json:"contender"`
			Oponent   int    `json:"oponent"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		battle, err := client.Battle.
			Create().SetResult(payload.Result).
			SetContenderID(payload.Contender).
			SetOponentID(payload.Oponent).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating battle: %w", err)
		}
		log.Println("battle created: ", battle)
		return c.Status(200).JSON(battle)
	})

	app.Get("/all/battle", func(c *fiber.Ctx) error {
		battles, err := client.Battle.Query().WithContender().WithContender().All(ctx)
		if err != nil {
			return fmt.Errorf("failed querying battles: %w", err)
		}
		log.Println("returned battles:", battles)
		return c.Status(200).JSON(battles)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.Get("APP_PORT"))))
}
