package features

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/log"
)

var _ api.MessageComponentFeature = (*ExampleFeature)(nil)

const (
	ExampleSelectCustomId = "example-select"
)

type ExampleFeature struct{}

func (h ExampleFeature) Names() []string {
	return []string{ExampleSelectCustomId}
}

func (h ExampleFeature) Handler(s *discordgo.Session, interaction *discordgo.Interaction) {
	if err := s.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello World!",
		},
	}); err != nil {
		log.Logger.Warnf("failed to respond to interaction: %v", err)
	}
}
