package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/goroutine/template/config"
	"github.com/goroutine/template/log"
)

var commands []SlashCommand

type ArgOption func(*SlashCommand)

type SlashCommand struct {
	Name                 string
	LocalizedName        *map[discordgo.Locale]string
	Description          string
	LocalizedDescription *map[discordgo.Locale]string
	Enabled              bool
	ArgsFunc             ArgOption
	AllowedRolesIds      []string
	ActionName           string
	args                 []SlashCommandArg
	Handler              func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type SlashCommandArg struct {
	Name                 string
	Description          string
	Type                 discordgo.ApplicationCommandOptionType
	Required             bool
	LocalizedName        *map[discordgo.Locale]string
	LocalizedDescription *map[discordgo.Locale]string
	Choices              []*discordgo.ApplicationCommandOptionChoice
}

func ArgsFromBuilders(args ...*SlashCommandArgBuilder) ArgOption {
	return func(c *SlashCommand) {
		for _, arg := range args {
			c.args = append(c.args, arg.Build())
		}
	}
}

func ArgsFromStructs(args ...SlashCommandArg) ArgOption {
	return func(c *SlashCommand) {
		c.args = append(c.args, args...)
	}
}

func GetCommands() []SlashCommand {
	return commands
}

func AddCommands(command ...SlashCommand) {
	commands = append(commands, command...)
}

func GetApplicationCommands() (applicationCommands []*discordgo.ApplicationCommand) {
	for _, command := range commands {
		var args []*discordgo.ApplicationCommandOption
		if command.ArgsFunc != nil {
			command.ArgsFunc(&command)
		}
		for _, arg := range command.args {
			var localizedName, localizedDescription map[discordgo.Locale]string
			if arg.LocalizedName != nil {
				localizedName = *arg.LocalizedName
			}
			if arg.LocalizedDescription != nil {
				localizedDescription = *arg.LocalizedDescription
			}
			args = append(args, &discordgo.ApplicationCommandOption{
				Name:                     arg.Name,
				NameLocalizations:        localizedName,
				DescriptionLocalizations: localizedDescription,
				Description:              arg.Description,
				Required:                 arg.Required,
				Choices:                  arg.Choices,
				Type:                     arg.Type,
			})
		}

		applicationCommands = append(applicationCommands, &discordgo.ApplicationCommand{
			Name:                     command.Name,
			NameLocalizations:        command.LocalizedName,
			Description:              command.Description,
			DescriptionLocalizations: command.LocalizedDescription,
			Options:                  args,
		})
	}
	return applicationCommands
}

func RegisterCommands(s *discordgo.Session, user *discordgo.User) {
	if guildId := config.ConfigInstance.GuildId; guildId != "" {
		for _, command := range commands {
			log.Logger.Debugf("Registering commands %s", command.Name)
		}
		if _, err := s.ApplicationCommandBulkOverwrite(user.ID, guildId, GetApplicationCommands()); err != nil {
			log.Logger.Error("Failed to register commands.", err)
			return
		}
		return
	}
	log.Logger.Warn("GuildId is not set. Commands will not be registered.")
}
