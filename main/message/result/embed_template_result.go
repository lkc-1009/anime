package result

type EmbedsResult struct {
}

type embedTemplateResult struct {
	Id              string       `json:"id"`
	ChannelId       string       `json:"channel_id"`
	GuildId         string       `json:"guild_id"`
	Timestamp       string       `json:"timestamp"`
	Tts             string       `json:"tts"`
	MentionEveryone string       `json:"mention_everyone"`
	Author          AuthorResult `json:"author"`
	Embeds          EmbedsResult `json:"embeds"`
	Pinned          string       `json:"pinned"`
	Type            string       `json:"type"`
	Flags           string       `json:"flags"`
}
