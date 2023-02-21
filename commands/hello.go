package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	helloCooldowns = make(map[string]time.Time)
)

func HelloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Vérifier si l'utilisateur a déjà utilisé la commande "hello" dans les 5 dernières secondes.
	if lastUsed, ok := helloCooldowns[m.Author.ID]; ok {
		if time.Since(lastUsed) < 5*time.Second {
			// L'utilisateur a utilisé la commande "hello" trop récemment. Envoyer un message d'erreur.
			s.ChannelMessageSend(m.ChannelID, "Vous devez attendre encore "+(5*time.Second-time.Since(lastUsed)).String()+" avant d'utiliser la commande de nouveau.")
			return
		}
	}

	// Enregistre l'heure actuelle comme dernière utilisation de la commande "hello".
	helloCooldowns[m.Author.ID] = time.Now()

	// Envoie un message "hello"
	s.ChannelMessageSend(m.ChannelID, "Bonjour !")
}
