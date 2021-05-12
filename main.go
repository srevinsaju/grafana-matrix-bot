package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/srevinsaju/grafana-matrix-bot/matrix"
	"github.com/srevinsaju/grafana-matrix-bot/types"
	"github.com/withmandala/go-log"
	"os"
)


var logger = log.New(os.Stdout)

// https://polyverse.com/blog/how-to-embed-versioning-information-in-go-applications-f76e2579b572/
var (
	BuildVersion string = ""
	BuildTime    string = ""
)


func main() {
	if BuildVersion != "" || BuildTime != "" {
		logger.Infof("Grafana Matrix Bot %s Build:%s", BuildVersion, BuildTime)
	} else {
		logger.Info("Grafana Matrix Bot (local dev build)")
	}
	app := fiber.New()
	logger.Infof("Starting fiber http server")

	lastArg := os.Args[len(os.Args)-1]
	if lastArg == "grafana-matrix-bot" {
		// the user has not provided any commands along with the executable name
		// so, we should show the usage
		logger.Info("grafana-matrix-bot : A grafana matrix bot to report spikes in statistics")
		logger.Info("")
		logger.Info("To load an existing configuration: ")
		logger.Info("  $ grafana-matrix-bot path/to/config.json")
		logger.Info("")
		return
	}

	if _, err := os.Stat(lastArg); os.IsNotExist(err) {
		logger.Fatal("The specified path does not exist")
	}

	cfg, err := matrix.ConfigFromFile(lastArg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("Connecting to matrix homeserver.")
	client := matrix.Setup(cfg)

	app.Post(fmt.Sprintf("/grafana/%s", cfg.SecretKey), func(c *fiber.Ctx) error {
		p := new(types.GrafanaWebhook)

		if err := c.BodyParser(p); err != nil {
			return err
		}
		_ = matrix.SendPhoto(client, cfg.ChannelId, p)
		err = matrix.SendMessage(client, cfg.ChannelId, p)
		if err != nil {
			return err
		}
		return c.SendStatus(200)

	})

	err = app.Listen(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Fatal(err)
	}
}
