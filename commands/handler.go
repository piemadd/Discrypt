package commands

import (
	"github.com/StephenSulimani/Discrypt/constants"
	"github.com/StephenSulimani/Discrypt/structs"
	"github.com/StephenSulimani/Discrypt/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// credit to https://github.com/apiks/ZeroTsu

var CommandMap = make(map[string]*structs.Command)
var AliasMap = make(map[string]string)

func RegisterCommand(c *structs.Command) {
	CommandMap[c.Trigger] = c
	for _, a := range c.Aliases {
		AliasMap[a] = c.Trigger
	}

	log.Println(c.Trigger + " loaded!")
}

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m == nil {
		return
	}
	if m.Author == nil {
		return
	}
	if m.Author.Bot {
		return
	}
	if m.Author.ID == "" {
		return
	}
	if m.Message == nil {
		return
	}
	if m.Message.Content == "" {
		return
	}

	if len(m.Message.Content) <= 1 || m.Message.Content[0:len(constants.PREFIX)] != constants.PREFIX {
		return
	}

	cmdTrigger := strings.Split(m.Content, " ")[0][len(constants.PREFIX):]
	cmdTrigger = strings.ToLower(cmdTrigger)
	cmd, ok := CommandMap[cmdTrigger]
	if !ok {
		cmd, ok = CommandMap[AliasMap[cmdTrigger]]
		if !ok {
			return
		}
	}

	if cmd.AdminOnly && !utils.StringInArray(m.Message.Author.ID, constants.DEV_IDS) {
		_, _ = s.ChannelMessageSend(m.Message.ChannelID, "You do not have permission to perform this command!")
		return
	}

	args := strings.Split(m.Message.Content, " ")
	copy(args[0:], args[1:])
	args[len(args)-1] = ""
	args = args[:len(args)-1]

	_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
	cmd.Exec(s, m.Message, args)
}
