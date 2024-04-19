package auto_role

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/embed"
	i18n "github.com/kaysoro/discordgo-i18n"
)

var valorantRolesIcons = []string{":unranked:1225196264929366187", ":iron_rank:1225168211037917307", ":bronze_rank:1225168255610519664", ":silver_rank:1225168292738760809", ":gold_rank:1225168329120157747",
	":platinum_rank:1225168355854520341", ":diamond_rank:1225168433931489290", ":ascendant_rank:1225168499354112040",
	":immortal_rank:1225168468496748645", ":radiant_rank:1225168530002018436"}

func CreateAutoRoleRankedEmbed(s *discordgo.Session) {
	autoRoleRankedChannel, err := s.Channel(config.ConfigInstance.Channels.AutoRoleRankedChannel)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	channelMessages, err := s.ChannelMessages(autoRoleRankedChannel.ID, 100, "", "", "")
	if err != nil {
		log.Logger.Error(err)
		return
	}

	if len(channelMessages) == 1 {
		return
	}

	var messagesToDelete []string
	for _, message := range channelMessages {
		messagesToDelete = append(messagesToDelete, message.ID)
	}

	err = s.ChannelMessagesBulkDelete(autoRoleRankedChannel.ID, messagesToDelete)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	autoRoleRankedMessage, err := s.ChannelMessageSendEmbeds(autoRoleRankedChannel.ID, embed.New().
		SetTitle(i18n.Get(discordgo.French, "auto_role_ranked.embed_title", nil)).
		SetDescription(i18n.Get(discordgo.French, "auto_role_ranked.embed_description")).
		SetColor(embed.VALOROUS).
		SetDefaultFooter().
		SetCurrentTimestamp().
		SetThumbnail("https://zupimages.net/up/24/14/malz.png").
		ToMessageEmbeds())
	if err != nil {
		log.Logger.Error(err)
		return
	}

	for _, icon := range valorantRolesIcons {
		err = s.MessageReactionAdd(autoRoleRankedChannel.ID, autoRoleRankedMessage.ID, icon)
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}
}
