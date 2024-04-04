package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/features/clip"
	"github.com/goroutine/template/features/suggestion"
	"github.com/rs/zerolog/log"
)

func MessageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.ChannelID {
	case config.ConfigInstance.Channels.SuggestionChannel:
		err := suggestion.SuggestionFeature(s, m)
		if err != nil {
			log.Logger.Err(err)
			return
		}

	case config.ConfigInstance.Channels.ClipChannel:
		err := clip.ClipFeature(s, m)
		if err != nil {
			log.Logger.Err(err)
			return
		}

	default:
		return
	}
}
