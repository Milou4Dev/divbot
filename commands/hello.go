package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

const helloCooldown = 5 * time.Second

var (
	helloCooldowns = make(map[string]time.Time)
)

func HelloCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Member.User.ID

	if lastUsed, ok := helloCooldowns[userID]; ok {
		if time.Since(lastUsed) < helloCooldown {
			remainingCooldown := helloCooldown - time.Since(lastUsed)
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

	helloCooldowns[userID] = time.Now()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseTypeChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello!",
		},
	})
}
