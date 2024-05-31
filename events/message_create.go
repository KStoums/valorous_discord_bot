package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/features/announce"
	"github.com/goroutine/template/features/changelog"
	"github.com/goroutine/template/features/clip"
	"github.com/goroutine/template/features/suggestion"
	"github.com/goroutine/template/log"
)

func MessageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.ChannelID {
	case config.ConfigInstance.Channels.SuggestionChannel:
		err := suggestion.SuggestionFeature(s, m)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		return

	case config.ConfigInstance.Channels.ClipChannel:
		err := clip.ClipFeature(s, m)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		return

	case config.ConfigInstance.Channels.PublicAnnounceChannel, config.ConfigInstance.Channels.TeamAnnounceChannel:
		err := announce.AnnounceFeature(s, m)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		return

	case config.ConfigInstance.Channels.BotChangelog:
		err := changelog.ChangelogFeature(s, m)
		if err != nil {
			log.Logger.Error(err)
			return
		}
		return

	default:
		return
	}
}
