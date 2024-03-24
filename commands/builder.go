package commands

import "github.com/bwmarrin/discordgo"

type SlashCommandArgBuilder struct {
	arg SlashCommandArg
}

func NewSlashCommandArgBuilder() *SlashCommandArgBuilder {
	return &SlashCommandArgBuilder{}
}

func (s *SlashCommandArgBuilder) SetName(name string) *SlashCommandArgBuilder {
	s.arg.Name = name
	return s
}

func (s *SlashCommandArgBuilder) SetLocalizedName(localizedName *map[discordgo.Locale]string) *SlashCommandArgBuilder {
	s.arg.LocalizedName = localizedName
	return s
}

func (s *SlashCommandArgBuilder) SetDescription(description string) *SlashCommandArgBuilder {
	s.arg.Description = description
	return s
}

func (s *SlashCommandArgBuilder) SetLocalizedDescription(localizedDescription *map[discordgo.Locale]string) *SlashCommandArgBuilder {
	s.arg.LocalizedDescription = localizedDescription
	return s
}

func (s *SlashCommandArgBuilder) SetRequired(required bool) *SlashCommandArgBuilder {
	s.arg.Required = required
	return s
}

func (s *SlashCommandArgBuilder) SetType(argType discordgo.ApplicationCommandOptionType) *SlashCommandArgBuilder {
	s.arg.Type = argType
	return s
}

func (s *SlashCommandArgBuilder) SetChoices(choices []*discordgo.ApplicationCommandOptionChoice) *SlashCommandArgBuilder {
	s.arg.Choices = choices
	return s
}

func (s *SlashCommandArgBuilder) AddStringChoice(name, value string) *SlashCommandArgBuilder {
	s.arg.Choices = append(s.arg.Choices, &discordgo.ApplicationCommandOptionChoice{
		Name:  name,
		Value: value,
	})
	return s
}

func (s *SlashCommandArgBuilder) AddIntChoice(name string, value int) *SlashCommandArgBuilder {
	s.arg.Choices = append(s.arg.Choices, &discordgo.ApplicationCommandOptionChoice{
		Name:  name,
		Value: value,
	})
	return s
}

func (s *SlashCommandArgBuilder) AddBoolChoice(name string, value bool) *SlashCommandArgBuilder {
	s.arg.Choices = append(s.arg.Choices, &discordgo.ApplicationCommandOptionChoice{
		Name:  name,
		Value: value,
	})
	return s
}

func (s *SlashCommandArgBuilder) Build() SlashCommandArg {
	return s.arg
}
