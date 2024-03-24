package config

// Config Define config here with substructs for each config category
type Config struct {
	GuildId  string   `json:"guildId"`
	Channels channels `json:"channels"`
}

type channels struct {
	MemberCount string `json:"memberCount"`
}
