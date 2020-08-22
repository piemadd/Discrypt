package commands

import (
	"github.com/StephenSulimani/Discrypt/constants"
	"github.com/StephenSulimani/Discrypt/structs"
	"github.com/StephenSulimani/Discrypt/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

func pingCommand(s *discordgo.Session, m *discordgo.Message, args []string) {
	embed := utils.NewEmbed().
		SetDescription("Ping...").
		SetColor(constants.HEX_COLOR).
		SetTimestamp(time.Now().Format(time.RFC3339)).
		SetAuthor(m.Author.Username + "#" + m.Author.Discriminator, m.Author.AvatarURL("1024"))
	start := time.Now()
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	if utils.NotifyError(err, s, m) { return }
	end := time.Now()

	diff := end.Sub(start)
	embed.SetDescription("Pong! Latency: "+strconv.Itoa(int(diff.Milliseconds()))+"ms")
	_, err = s.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, embed.MessageEmbed)

	_ = utils.NotifyError(err, s, m)
}

func init() {
	RegisterCommand(&structs.Command{
		Exec:      pingCommand,
		Trigger:   "ping",
		Aliases:   nil,
		Usage:     "ping",
		Desc:      "Get latency",
		AdminOnly: false,
	})
}
