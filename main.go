package main

import (
	"os"
	"strconv"

	"github.com/bold-commerce/go-shopify"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Error(err)
		return
	}

	debug := GetBool("APP_DEBUG", false)
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Error(err)
		return
	}

	shopifyClient := goshopify.NewClient(goshopify.App{
		ApiKey:   os.Getenv("APP_KEY"),
		Password: os.Getenv("APP_PASSWORD"),
	}, os.Getenv("SHOP_NAME"), "")

	log.Info("Bot started!")
	err = NewHandler(shopifyClient, bot).ListenAndServe()
	if err != nil {
		log.Error(err)
	}
}

// GetBool reads env and parses
func GetBool(env string, defaultVal bool) bool {
	val, err := strconv.ParseBool(os.Getenv(env))
	if err != nil {
		return defaultVal
	}
	return val
}
