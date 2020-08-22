package structs

import "github.com/bwmarrin/discordgo"

type Command struct {
	Exec       func(*discordgo.Session, *discordgo.Message, []string)
	Trigger    string
	Aliases    []string
	Usage      string
	Desc       string
	AdminOnly  bool
}
