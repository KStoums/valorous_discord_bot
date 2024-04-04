package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/features/suggestion"
)

func MessageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID != config.ConfigInstance.Channels.SuggestionChannel {
		return
	}

	suggestion.SuggestionFeature(s, m)
}
