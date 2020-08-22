package utils

import (
	"github.com/StephenSulimani/Discrypt/constants"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandleFatal(loc string, err error) {
	if err != nil {
		log.Fatal("Fatal error at " + loc + "\n\n" + err.Error())
	}
}

func HandleError(err error, s *discordgo.Session) bool {
	if err != nil {
		log.Println("An error has occurred: ", err.Error())

		embed := NewEmbed().
			SetTitle("Error").
			SetDescription("```" + err.Error() + "```").
			SetColor(0x992D22).
			SetTimestamp(time.Now().Format(time.RFC3339)).MessageEmbed
		comp := &discordgo.MessageSend{
			Content: "@everyone",
			Embed:   embed,
		}
		_, _ = s.ChannelMessageSendComplex(constants.ALERTS, comp)

		return true
	}

	return false
}

func NotifyError(err error, s *discordgo.Session, m *discordgo.Message) bool {
	if err != nil {
		log.Println("An error has occurred: ", err.Error())

		embed := NewEmbed().
			SetTitle("Error").
			SetDescription("```" + err.Error() + "```").
			SetColor(0x992D22).
			SetTimestamp(time.Now().Format(time.RFC3339)).MessageEmbed
		comp := &discordgo.MessageSend{
			Content: "@everyone",
			Embed:   embed,
		}
		_, _ = s.ChannelMessageSendComplex(constants.ALERTS, comp)

		embed = NewEmbed().
			SetTitle("Error!").
			SetColor(0x992D22).
			SetTimestamp(time.Now().Format(time.RFC3339)).
			SetDescription("An error occurred! Please try again later.").MessageEmbed
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)

		return true
	}

	return false
}

func StringInArray(query string, array []string) bool {
	for _, x := range array {
		if query == x {
			return true
		}
	}

	return false
}
