package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

const pingCooldown = 5 * time.Second

var (
	pingCooldowns = make(map[string]time.Time)
)

func PingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Member.User.ID

	if lastUsed, ok := pingCooldowns[userID]; ok {
		if time.Since(lastUsed) < pingCooldown {
			remainingCooldown := pingCooldown - time.Since(lastUsed)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseTypeChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("You need to wait %s before using this command again.", remainingCooldown),
					Flags:   64, // Set the message to be ephemeral (only visible to the user)
				},
			})
			return
		}
	}

	pingCooldowns[userID] = time.Now()

	apiPing := s.HeartbeatLatency().Round(time.Millisecond)

	embed := &discordgo.MessageEmbed{
		Title:       "Bot's speed:",
		Description: fmt.Sprintf("API (Discord) Ping: %s\nBot Ping: %s", apiPing, s.HeartbeatLatency()),
		Color:       27157218255, // Discord color
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseTypeChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
