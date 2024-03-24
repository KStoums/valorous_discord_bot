package helpers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
)

func GetMembersCount(s *discordgo.Session) (int, error) {
	var membersCount int
	var lastMemberId string

	for {
		members, err := s.GuildMembers(config.ConfigInstance.GuildId, lastMemberId, 1000)
		if err != nil {
			return 0, err
		}
		membersCount += len(members)
		if len(members) < 1000 {
			break
		}

		lastMemberId = members[len(members)-1].User.ID
	}

	return membersCount, nil
}
