package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ReadyEvent(s *discordgo.Session, e *discordgo.Ready) {
	log.Println("Bot is ready to go!\n\nLogged in as: " + e.User.Username + "#" + e.User.Discriminator)
	s.UpdateListeningStatus("status lol")
}
