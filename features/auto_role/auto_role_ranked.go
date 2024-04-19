package auto_role

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/strutils"
	"strings"
)

func AddAutoRoleRankedFeature(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	roles, err := s.GuildRoles(r.GuildID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	emojiNameSplit := strings.Split(r.Emoji.Name, "_")
	if len(emojiNameSplit) != 2 {
		log.Logger.Error(errors.New("emojiNameSplit len is not equal as 2"))
		return
	}

	rankName := emojiNameSplit[0]

	for _, role := range roles {
		if strings.EqualFold(role.Name, rankName) {
			err = checkPlayerHasRankRole(s, r)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			err = s.GuildMemberRoleAdd(r.GuildID, r.Member.User.ID, role.ID)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			err = s.GuildMemberRoleAdd(r.GuildID, r.Member.User.ID, config.ConfigInstance.Roles.RankSeparatorRole)
			if err != nil {
				log.Logger.Error(err)
				return
			}
			return
		}
	}

	log.Logger.Error(errors.New("role not found by rank name"))
}

func RemoveAutoRoleRankedFeature(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	roles, err := s.GuildRoles(r.GuildID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	emojiNameSplit := strings.Split(r.Emoji.Name, "_")
	if len(emojiNameSplit) != 2 {
		log.Logger.Error(errors.New("emojiNameSplit len is not equal as 2"))
		return
	}
	rankName := emojiNameSplit[0]

	member, err := s.GuildMember(r.GuildID, r.UserID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	for _, role := range roles {
		if strings.EqualFold(role.Name, rankName) && strutils.ContainString(member.Roles, role.ID) {
			err = s.GuildMemberRoleRemove(r.GuildID, r.UserID, role.ID)
			if err != nil {
				log.Logger.Error(err)
				return
			}

			err = s.GuildMemberRoleRemove(r.GuildID, r.UserID, config.ConfigInstance.Roles.RankSeparatorRole)
			if err != nil {
				log.Logger.Error(err)
				return
			}
			return
		}
	}

	log.Logger.Error(errors.New("role not found by rank name"))
}

func checkPlayerHasRankRole(s *discordgo.Session, r *discordgo.MessageReactionAdd) error {
	member, err := s.GuildMember(r.GuildID, r.Member.User.ID)
	if err != nil {
		return err
	}

	for _, role := range member.Roles {
		if role == config.ConfigInstance.Roles.UnrankedRole || role == config.ConfigInstance.Roles.IronRole || role == config.ConfigInstance.Roles.BronzeRole ||
			role == config.ConfigInstance.Roles.SilverRole || role == config.ConfigInstance.Roles.GoldRole || role == config.ConfigInstance.Roles.PlatinumRole ||
			role == config.ConfigInstance.Roles.DiamondRole || role == config.ConfigInstance.Roles.AscendantRole || role == config.ConfigInstance.Roles.ImmortalRole ||
			role == config.ConfigInstance.Roles.RadiantRole {
			return errors.New("member has already a rank role")
		}
	}

	return nil
}
