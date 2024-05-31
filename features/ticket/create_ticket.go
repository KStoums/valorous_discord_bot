package ticket

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
	"strings"
)

const (
	createTicket  = "create-ticket"
	newTicketName = "ticket-"
)

var _ api.MessageComponentFeature = (*CreateTicketFeature)(nil)

type CreateTicketFeature struct{}

func (c *CreateTicketFeature) Names() []string {
	return []string{createTicket}
}

func (c *CreateTicketFeature) Handler(s *discordgo.Session, i *discordgo.Interaction) {
	if !checkIfMemberHasAlreadyTicket(s, i) {
		err := s.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: embed.New().
					SetTitle(i18n.Get(discordgo.French, "ticket.already_have_ticket_title")).
					SetDescription(i18n.Get(discordgo.French, "ticket.already_have_ticket_description")).
					SetColor(embed.VALOROUS).
					SetCurrentTimestamp().
					SetThumbnail("https://zupimages.net/up/24/22/4vnp.png").
					SetDefaultFooter().
					ToMessageEmbeds(),
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			log.Logger.Error(err)
			return
		}
		return
	}

	ticketChannel, err := s.GuildChannelCreateComplex(config.ConfigInstance.GuildId, discordgo.GuildChannelCreateData{
		Name:     newTicketName + i.Member.User.GlobalName,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: config.ConfigInstance.Channels.TicketOpenCategory,
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}

	ticketChannel.PermissionOverwrites = append(ticketChannel.PermissionOverwrites, &discordgo.PermissionOverwrite{
		ID:    i.Member.User.ID,
		Type:  discordgo.PermissionOverwriteTypeMember,
		Allow: discordgo.PermissionViewChannel,
	})

	_, err = s.ChannelEditComplex(ticketChannel.ID, &discordgo.ChannelEdit{
		PermissionOverwrites: ticketChannel.PermissionOverwrites,
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}

	_, err = s.ChannelMessageSendComplex(ticketChannel.ID, &discordgo.MessageSend{
		Embeds: embed.New().
			SetTitle(i18n.Get(discordgo.French, "ticket.created_ticket_title", i18n.Vars{
				"ticketAuthor": i.Member.User.GlobalName,
			})).
			SetDescription(i18n.Get(discordgo.French, "ticket.created_ticket_description")).
			SetColor(embed.VALOROUS).
			SetCurrentTimestamp().
			SetDefaultFooter().
			SetThumbnail(i.Member.User.AvatarURL("")).
			ToMessageEmbeds(),
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Label: "Clôturer le ticket",
						Style: discordgo.DangerButton,
						Emoji: &discordgo.ComponentEmoji{
							Name: "🔐",
						},
						CustomID: closeTicket,
					},
				},
			},
		},
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed.New().
				SetTitle(i18n.Get(discordgo.French, "ticket.created_ticket_title_response")).
				SetDescription(i18n.Get(discordgo.French, "ticket.created_ticket_description_response", i18n.Vars{
					"ticketChannelMention": ticketChannel.Mention(),
				})).
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetThumbnail("https://zupimages.net/up/24/22/vr0y.png").
				ToMessageEmbeds(),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = utils.SendLogToDiscordLogChannel(s, i18n.Get(discordgo.French, "ticket.ticket_created_logs_description",
		i18n.Vars{
			"ticketAuthor":         i.Member.Mention(),
			"ticketChannelMention": ticketChannel.Mention(),
		}))
	if err != nil {
		log.Logger.Error(err)
	}
}

func checkIfMemberHasAlreadyTicket(s *discordgo.Session, i *discordgo.Interaction) bool {
	guildChannels, err := s.GuildChannels(config.ConfigInstance.GuildId)
	if err != nil {
		log.Logger.Error(err)
		return false
	}

	for _, channel := range guildChannels {
		if channel.Name == newTicketName+strings.ToLower(i.Member.User.GlobalName) {
			return false
		}
	}

	return true
}
