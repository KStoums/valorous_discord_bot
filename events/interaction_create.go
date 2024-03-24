package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/commands"
	features2 "github.com/goroutine/template/features"
	"strings"
)

var features []api.MessageComponentFeature

func init() {
	features = []api.MessageComponentFeature{
		features2.ExampleFeature{},
	}
}

func InteractionCreateEvent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		for _, command := range commands.GetCommands() {
			if command.Name == i.ApplicationCommandData().Name {
				command.Handler(s, i)
				return
			}
		}
		break
	case discordgo.InteractionMessageComponent:
		data := i.MessageComponentData()
		for _, feature := range features {
			for _, name := range feature.Names() {
				if name == data.CustomID || strings.HasPrefix(data.CustomID, name) {
					feature.Handler(s, i.Interaction)
					return
				}
			}
		}
	}
}
