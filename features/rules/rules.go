package rules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/embed"
	"github.com/goroutine/template/utils/strutils"
	i18n "github.com/kaysoro/discordgo-i18n"
)

var _ api.MessageComponentFeature = (*RulesFeature)(nil)

const (
	AcceptRules = "accept-rules"
)

type RulesFeature struct{}

func (h *RulesFeature) Names() []string {
	return []string{AcceptRules}
}

func (h *RulesFeature) Handler(s *discordgo.Session, i *discordgo.Interaction) {
	member, err := s.GuildMember(i.GuildID, i.Member.User.ID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	if strutils.ContainString(member.Roles, config.ConfigInstance.Roles.AcceptedRuleRole) {
		err := s.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: embed.New().
					SetTitle(i18n.Get(discordgo.French, "rules.error.already_accepted_rules_title")).
					SetDescription(i18n.Get(discordgo.French, "rules.error.already_accepted_rules_description")).
					SetColor(embed.VALOROUS).
					SetCurrentTimestamp().
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

	err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, config.ConfigInstance.Roles.AcceptedRuleRole)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed.New().
				SetTitle(i18n.Get(discordgo.French, "rules.embed.accepted_rules_title")).
				SetDescription(i18n.Get(discordgo.French, "rules.embed.accepted_rules_description")).
				SetCurrentTimestamp().
				SetDefaultFooter().
				SetColor(embed.VALOROUS).
				ToMessageEmbeds(),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
