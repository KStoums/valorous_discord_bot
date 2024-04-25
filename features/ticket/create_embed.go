package ticket

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func CreateTicketEmbed(s *discordgo.Session) {
	ticketChannelMessage, err := s.ChannelMessages(config.ConfigInstance.Channels.TicketChannel, 100, "", "", "")
	if err != nil {
		log.Logger.Error(err)
		return
	}

	if len(ticketChannelMessage) == 1 {
		return
	} else if len(ticketChannelMessage) > 1 {
		var messagesToDelete []string
		for _, message := range ticketChannelMessage {
			messagesToDelete = append(messagesToDelete, message.ID)
		}
		err = s.ChannelMessagesBulkDelete(config.ConfigInstance.Channels.TicketChannel, messagesToDelete)
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}

	_, err = s.ChannelMessageSendComplex(config.ConfigInstance.Channels.TicketChannel, &discordgo.MessageSend{
		Embeds: embed.New().
			SetTitle(i18n.Get(discordgo.French, "ticket.embed_title")).
			SetDescription(i18n.Get(discordgo.French, "ticket.embed_description")).
			SetColor(embed.VALOROUS).
			SetCurrentTimestamp().
			SetDefaultFooter().
			SetThumbnail("https://zupimages.net/up/24/16/itry.png").
			ToMessageEmbeds(),
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: i18n.Get(discordgo.French, "ticket.button.open_ticket"),
						Style: discordgo.SecondaryButton,
						Emoji: &discordgo.ComponentEmoji{
							Name: "🎟️",
						},
						CustomID: createTicket,
					},
				},
			},
		},
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
