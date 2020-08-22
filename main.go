package main

import (
	_ "fmt"
	"github.com/StephenSulimani/Discrypt/commands"
	"os"
	"os/signal"
	"syscall"

	"github.com/StephenSulimani/Discrypt/events"
	"github.com/StephenSulimani/Discrypt/utils"
	"github.com/bwmarrin/discordgo"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	utils.HandleFatal("Creating DG session", err)

	dg.AddHandler(events.ReadyEvent)
	dg.AddHandler(commands.HandleCommand)

	err = dg.Open()
	utils.HandleFatal("Opening DG session", err)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()
}
