package clip

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"strings"
)

func ClipFeature(s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if err != nil {
		return err
	}

	if !strings.Contains(m.Message.Content, "https://") {
		return err
	}

	var clipPlatform string
	if strings.Contains(m.Message.Content, "youtube") {
		clipPlatform = "Youtube"
	} else if strings.Contains(m.Message.Content, "twitch") {
		clipPlatform = "Twitch"
	}

	clipEmbed, err := s.ChannelMessageSendEmbed(config.ConfigInstance.Channels.ClipChannel, embed.New().
		SetTitle(i18n.Get(discordgo.French, "clip.clip_title", i18n.Vars{"authorName": m.Author.GlobalName})).
		AddInlinedField("🎬 Plateforme", clipPlatform).
		AddInlinedField("🔗 URL", "[Cliquez-ici]("+m.Message.Content+")").
		SetColor(embed.VALOROUS).
		SetDefaultFooter().
		SetCurrentTimestamp().
		SetThumbnail("https://zupimages.net/up/24/22/79cm.png").
		SetVideo(m.Message.Content).
		SetImage(m.Message.Embeds[0].Thumbnail.URL).
		ToMessageEmbed(),
	)
	if err != nil {
		return err
	}

	_, err = s.MessageThreadStart(m.ChannelID, clipEmbed.ID, i18n.Get(discordgo.French, "clip.clip_title", i18n.Vars{"authorName": m.Author.GlobalName}), 0)
	if err != nil {
		return err
	}
	return nil
}
