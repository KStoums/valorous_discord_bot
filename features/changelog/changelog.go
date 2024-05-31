package changelog

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func ChangelogFeature(s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed.New().
		SetTitle(i18n.Get(discordgo.French, "changelog.changelog_title")).
		SetDescription(m.Message.Content).
		SetThumbnail("https://zupimages.net/up/24/22/e0rz.png").
		SetCurrentTimestamp().
		SetDefaultFooter().
		SetColor(embed.VALOROUS).
		ToMessageEmbed())
	return err
}
