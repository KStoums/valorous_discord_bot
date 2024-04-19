package ticket

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"time"
)

const closeTicket = "close-ticket"

var _ api.MessageComponentFeature = (*CloseTicketFeature)(nil)

type CloseTicketFeature struct{}

func (c *CloseTicketFeature) Names() []string {
	return []string{closeTicket}
}

func (c *CloseTicketFeature) Handler(s *discordgo.Session, i *discordgo.Interaction) {
	ticketChannel, err := s.Channel(i.ChannelID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	if ticketChannel.ParentID != config.ConfigInstance.Channels.TicketOpenCategory {
		return
	}

	err = s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed.New().
				SetTitle(i18n.Get(discordgo.French, "ticket.close_ticket_title")).
				SetDescription(i18n.Get(discordgo.French, "ticket.close_ticket_description")).
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetColor(embed.VALOROUS).
				ToMessageEmbeds(),
		},
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}

	time.Sleep(5 * time.Second)
	_, err = s.ChannelDelete(ticketChannel.ID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = utils.SendLogToDiscordLogChannel(s, i18n.Get(discordgo.French, "ticket.closed_ticket_logs_description", i18n.Vars{
		"member":            i.Member.Mention(),
		"ticketChannelName": ticketChannel.Name,
	}))
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
