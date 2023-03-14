package commands

import (
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Data.Name {
	case "hello":
		HelloCommand(s, i)
	case "ping":
		PingCommand(s, i)
	}
}
