package template

import "github.com/tencent-connect/botgo/dto"

// CreateEmbedTemplate 创建embed消息
// Content-Type -> application/json
// obj -> {{Name: "text",}, {}...}
func createEmbedTemplate(
	tips string,
	title string,
	picUrl string,
	obj []*dto.EmbedField,
) *dto.Embed {
	return &dto.Embed{
		Title: title,
		Prompt: tips,
		Thumbnail: dto.MessageEmbedThumbnail{URL: picUrl},
		Fields: obj,
	}
}


