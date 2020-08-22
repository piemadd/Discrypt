package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func ReadyEvent(_ *discordgo.Session, e *discordgo.Ready) {
	log.Println("Bot is ready to go!\n\nLogged in as: " + e.User.Username + "#" + e.User.Discriminator)
}
