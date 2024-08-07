package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/FastLane-Labs/fastlane-online-solver/bot"
	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/datagen"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
)

func main() {
	config := config.NewConfig()
	log.InitLogger(config.LogLevel)

	switch config.Mode {
	case "run_bot":
		runBot(config)
	case "data_gen":
		if err := datagen.DataGen(config); err != nil {
			log.Error("Failed to generate data files", "error", err)
			os.Exit(1)
		}
	default:
		log.Error("Invalid mode", "mode", config.Mode, "acceptedModes", "run_bot/data_gen")
	}
}

func runBot(config *config.Config) {
	bot, err := bot.NewBot(config)
	if err != nil {
		panic("Failed to create bot - " + err.Error())
	}
	bot.Run()

	var shutdown = make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	for range shutdown {
		log.Info("Shutting down bot...")
		os.Exit(1)
	}
}
