package rule

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func CreateRuleEmbed(s *discordgo.Session) {
	channelMessages, err := s.ChannelMessages(config.ConfigInstance.Channels.RulesChannel, 100, "", "", "")
	if err != nil {
		log.Logger.Error(err)
		return
	}

	if len(channelMessages) != 1 {
		var messageList []string
		for _, message := range channelMessages {
			messageList = append(messageList, message.ID)
		}

		err = s.ChannelMessagesBulkDelete(config.ConfigInstance.Channels.RulesChannel, messageList)
		if err != nil {
			log.Logger.Error(err)
			return
		}

		_, err = s.ChannelMessageSendComplex(config.ConfigInstance.Channels.RulesChannel, &discordgo.MessageSend{
			Embeds: embed.New().
				SetTitle(i18n.Get(discordgo.French, "rules.rules_title")).
				SetDescription(i18n.Get(discordgo.French, "rules.rules_description")).
				SetThumbnail("https://zupimages.net/up/24/22/79cm.png").
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetColor(embed.VALOROUS).
				ToMessageEmbeds(),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: i18n.Get(discordgo.French, "rules.button.accept_rules"),
							Style: discordgo.SuccessButton,
							Emoji: &discordgo.ComponentEmoji{
								Name: "🖊️",
							},
							CustomID: acceptRules,
						},
					},
				},
			},
		})
		if err != nil {
			log.Logger.Error(err)
			return
		}

		log.Logger.Warn("The regulation embed was not found, so it was created")
	}
}
