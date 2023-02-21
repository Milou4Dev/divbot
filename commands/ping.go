package commands

import (
    "time"

    "github.com/bwmarrin/discordgo"
)

var (
    pingCooldowns = make(map[string]time.Time)
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Vérifier si l'utilisateur a déjà utilisé la commande "ping" dans les 5 dernières secondes.
    if lastUsed, ok := pingCooldowns[m.Author.ID]; ok {
        if time.Now().Sub(lastUsed) < 5*time.Second {
            // L'utilisateur a utilisé la commande "ping" trop récemment. Envoyer un message d'erreur.
            s.ChannelMessageSend(m.ChannelID, "Vous devez attendre encore "+(5*time.Second-time.Now().Sub(lastUsed)).String()+" avant d'utiliser la commande de nouveau.")
            return
        }
    }

    // Enregistre l'heure actuelle comme dernière utilisation de la commande "ping".
    pingCooldowns[m.Author.ID] = time.Now()

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
