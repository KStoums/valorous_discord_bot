package api

import "github.com/bwmarrin/discordgo"

type MessageComponentFeature interface {
	Names() []string
	Handler(s *discordgo.Session, i *discordgo.Interaction)
}
