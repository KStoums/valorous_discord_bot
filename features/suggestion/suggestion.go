package suggestion

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func SuggestionFeature(s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if err != nil {
		return err
	}

	suggestionEmbed, err := s.ChannelMessageSendEmbed(config.ConfigInstance.Channels.SuggestionChannel, embed.New().
		SetTitle(i18n.Get(discordgo.French, "suggestion.suggestion_title", i18n.Vars{"authorName": m.Author.GlobalName})).
		SetDescription(m.Message.Content).
		SetColor(embed.VALOROUS).
		SetDefaultFooter().
		SetCurrentTimestamp().
		SetThumbnail("https://zupimages.net/up/24/22/oo4p.png").
		ToMessageEmbed())
	if err != nil {
		return err
	}

	_, err = s.MessageThreadStart(m.ChannelID, suggestionEmbed.ID, i18n.Get(discordgo.French, "suggestion.suggestion_title", i18n.Vars{"authorName": m.Author.GlobalName}), 0)
	if err != nil {
		return err
	}

	err = s.MessageReactionAdd(m.ChannelID, suggestionEmbed.ID, "✅")
	if err != nil {
		return err
	}

	return s.MessageReactionAdd(m.ChannelID, suggestionEmbed.ID, "❌")
}
