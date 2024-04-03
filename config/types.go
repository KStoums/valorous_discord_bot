package config

// Config Define config here with substructs for each config category
type Config struct {
	GuildId  string   `json:"guildId"`
	Channels channels `json:"channels"`
	Roles    roles    `json:"roles"`
}

type channels struct {
	MemberCount          string `json:"memberCount"`
	WelcomeChannel       string `json:"welcomeChannel"`
	RulesChannel         string `json:"rulesChannel"`
	AutoVoiceChannel     string `json:"autoVoiceChannel"`
	AutoVoiceChannelTeam string `json:"autoVoiceChannelTeam"`
	TeamCategory         string `json:"teamCategory"`
	AutoVoiceCategory    string `json:"autoVoiceCategory"`
}

type roles struct {
	MemberRole       string `json:"memberRole"`
	AcceptedRuleRole string `json:"acceptedRuleRole"`
}
