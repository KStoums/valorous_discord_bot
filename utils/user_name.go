package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func GetUserDisplayName(user *discordgo.User) string {
	if user.GlobalName != "" {
		return user.GlobalName
	}

	return fmt.Sprintf("%s#%s", user.Username, user.Discriminator)
}

func GetMemberDisplayNameWithNick(member *discordgo.Member) string {
	if member.Nick != "" {
		return member.Nick
	}

	return GetUserDisplayName(member.User)
}

func GetMemberDisplayName(member *discordgo.Member) string {
	return GetUserDisplayName(member.User)
}
