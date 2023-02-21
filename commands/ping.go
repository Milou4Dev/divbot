package commands

import (
    "time"

    "github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Enregistre l'heure actuelle
    start := time.Now()

    // Envoie un message "ping"
    msg, err := s.ChannelMessageSend(m.ChannelID, "ping")
    if err != nil {
        // En cas d'erreur, envoie un message d'erreur
        s.ChannelMessageSend(m.ChannelID, "Une erreur est survenue lors de l'envoi du message.")
        return
    }

    // Calcule la durée entre l'envoi du message et sa réception
    end := time.Now()
    duration := end.Sub(start)

    // Met à jour le message avec la durée
    s.ChannelMessageEdit(m.ChannelID, msg.ID, "Pong! Temps de réponse : "+duration.String())
}
