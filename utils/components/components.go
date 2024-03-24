package components

import "github.com/bwmarrin/discordgo"

// The goal os this file is to make it easier to create components for slash commands.
// The components are created with the help of the ComponentBuilder struct.

type ComponentBuilder struct {
	Components []discordgo.MessageComponent
}

type SelectMenuBuilder struct {
	selectMenu discordgo.SelectMenu
}

type SelectOptionBuilder struct {
	option discordgo.SelectMenuOption
}

type ActionRowBuilder struct {
	actionRow discordgo.ActionsRow
}

type ButtonBuilder struct {
	button discordgo.Button
}

// NewComponentBuilder creates a new ComponentBuilder.
func NewComponentBuilder() *ComponentBuilder {
	return &ComponentBuilder{}
}

// NewSelectMenuBuilder creates a new SelectMenuBuilder.
func NewSelectMenuBuilder() *SelectMenuBuilder {
	return &SelectMenuBuilder{}
}

// NewActionRowBuilder creates a new ActionRowBuilder.
func NewActionRowBuilder() *ActionRowBuilder {
	return &ActionRowBuilder{}
}

// NewButtonBuilder creates a new ButtonBuilder.
func NewButtonBuilder() *ButtonBuilder {
	return &ButtonBuilder{}
}

// NewSelectOptionBuilder creates a new SelectOptionBuilder.
func NewSelectOptionBuilder() *SelectOptionBuilder {
	return &SelectOptionBuilder{}
}

// AddSelectMenu adds a SelectMenu to the ComponentBuilder.
func (cb *ComponentBuilder) AddSelectMenu(sm *SelectMenuBuilder) {
	cb.Components = append(cb.Components, sm.Build())
}

// AddButton adds a Button to the ComponentBuilder.
func (cb *ComponentBuilder) AddButton(b *ButtonBuilder) *ComponentBuilder {
	cb.Components = append(cb.Components, b.Build())
	return cb
}

// AddActionRow adds an ActionRow to the ComponentBuilder.
func (cb *ComponentBuilder) AddActionRow(ar *ActionRowBuilder) *ComponentBuilder {
	cb.Components = append(cb.Components, ar.Build())
	return cb
}

// SetCustomId sets the CustomId of the SelectMenuBuilder.
func (sb *SelectMenuBuilder) SetCustomId(customId string) *SelectMenuBuilder {
	sb.selectMenu.CustomID = customId
	return sb
}

// SetMenuType sets the MenuType of the SelectMenuBuilder.
func (sb *SelectMenuBuilder) SetMenuType(menuType discordgo.SelectMenuType) *SelectMenuBuilder {
	sb.selectMenu.MenuType = menuType
	return sb
}

// SetPlaceholder sets the Placeholder of the SelectMenuBuilder.
func (sb *SelectMenuBuilder) SetPlaceholder(placeholder string) *SelectMenuBuilder {
	sb.selectMenu.Placeholder = placeholder
	return sb
}

// SetMinValues sets the MinValues of the SelectMenuBuilder.
func (sb *SelectMenuBuilder) SetMinValues(minValues int) *SelectMenuBuilder {
	sb.selectMenu.MinValues = &minValues
	return sb
}

// SetMaxValues sets the MaxValues of the SelectMenuBuilder.
func (sb *SelectMenuBuilder) SetMaxValues(maxValues int) *SelectMenuBuilder {
	sb.selectMenu.MaxValues = maxValues
	return sb
}

// SetDisabled sets the Disabled of the SelectMenuBuilder.
func (sb *SelectMenuBuilder) SetDisabled(disabled bool) *SelectMenuBuilder {
	sb.selectMenu.Disabled = disabled
	return sb
}

// AddOption adds an Option to the SelectMenuBuilder.
func (sb *SelectMenuBuilder) AddOption(option *SelectOptionBuilder) *SelectMenuBuilder {
	sb.selectMenu.Options = append(sb.selectMenu.Options, option.option)
	return sb
}

// AddOptions adds Options to the SelectMenuBuilder.
func (sb *SelectMenuBuilder) AddOptions(options ...*SelectOptionBuilder) *SelectMenuBuilder {
	for _, option := range options {
		sb.selectMenu.Options = append(sb.selectMenu.Options, option.option)
	}
	return sb
}

// SetLabel sets the Label of the SelectOptionBuilder.
func (ob *SelectOptionBuilder) SetLabel(label string) *SelectOptionBuilder {
	ob.option.Label = label
	return ob
}

// SetValue sets the Value of the SelectOptionBuilder.
func (ob *SelectOptionBuilder) SetValue(value string) *SelectOptionBuilder {
	ob.option.Value = value
	return ob
}

// SetDescription sets the Description of the SelectOptionBuilder.
func (ob *SelectOptionBuilder) SetDescription(description string) *SelectOptionBuilder {
	ob.option.Description = description
	return ob
}

// SetEmoji sets the Emoji of the SelectOptionBuilder.
func (ob *SelectOptionBuilder) SetEmoji(emoji discordgo.ComponentEmoji) *SelectOptionBuilder {
	ob.option.Emoji = emoji
	return ob
}

// SetDefault sets the Default of the SelectOptionBuilder.
func (ob *SelectOptionBuilder) SetDefault(defaultValue bool) *SelectOptionBuilder {
	ob.option.Default = defaultValue
	return ob
}

// AddSelectMenu adds a SelectMenu to the ActionRowBuilder.
func (ab *ActionRowBuilder) AddSelectMenu(sm *SelectMenuBuilder) *ActionRowBuilder {
	ab.actionRow.Components = append(ab.actionRow.Components, sm.Build())
	return ab
}

// AddButton adds a Button to the ActionRowBuilder.
func (ab *ActionRowBuilder) AddButton(b *ButtonBuilder) *ActionRowBuilder {
	ab.actionRow.Components = append(ab.actionRow.Components, b.Build())
	return ab
}

// AddButtons adds Buttons to the ActionRowBuilder.
func (ab *ActionRowBuilder) AddButtons(buttons ...*ButtonBuilder) *ActionRowBuilder {
	for _, button := range buttons {
		ab.actionRow.Components = append(ab.actionRow.Components, button.Build())
	}
	return ab
}

func (bb *ButtonBuilder) SetLabel(label string) *ButtonBuilder {
	bb.button.Label = label
	return bb
}

func (bb *ButtonBuilder) SetStyle(style discordgo.ButtonStyle) *ButtonBuilder {
	bb.button.Style = style
	return bb
}

func (bb *ButtonBuilder) SetEmoji(emoji discordgo.ComponentEmoji) *ButtonBuilder {
	bb.button.Emoji = emoji
	return bb
}

func (bb *ButtonBuilder) SetCustomId(customID string) *ButtonBuilder {
	bb.button.CustomID = customID
	return bb
}

func (bb *ButtonBuilder) SetURL(url string) *ButtonBuilder {
	bb.button.URL = url
	return bb
}

func (bb *ButtonBuilder) SetDisabled(disabled bool) *ButtonBuilder {
	bb.button.Disabled = disabled
	return bb
}

// Build builds the ComponentBuilder and returns the components.
func (cb *ComponentBuilder) Build() []discordgo.MessageComponent {
	return cb.Components
}

// Build builds the SelectMenuBuilder and returns the components.
func (sb *SelectMenuBuilder) Build() discordgo.SelectMenu {
	return sb.selectMenu
}

// Build builds the ActionRowBuilder and returns the components.
func (ab *ActionRowBuilder) Build() discordgo.ActionsRow {
	return ab.actionRow
}

// Build builds the ButtonBuilder and returns the components.
func (bb *ButtonBuilder) Build() discordgo.Button {
	return bb.button
}
