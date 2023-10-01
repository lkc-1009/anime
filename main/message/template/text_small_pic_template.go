package template

import "github.com/tencent-connect/botgo/dto"

// CreateTextSmallPicTemplate 创建文本+缩略图模板
func createTextSmallPicTemplate(
	desc string,
	tips string,
	title string,
	titleDesc string,
	picUrl string,
	redirectUrl string,
	footerText string,
) *dto.Ark {
	return &dto.Ark{
		TemplateID: 24,
		KV: []*dto.ArkKV{
			{
				Key:   "#DESC#",
				Value: desc,
			},
			{
				Key:   "#PROMPT#",
				Value: tips,
			},
			{
				Key:   "#TITLE#",
				Value: title,
			},
			{
				Key:   "#METADESC#",
				Value: titleDesc,
			},
			{
				Key:   "#IMG#",
				Value: picUrl,
			},
			{
				Key:   "#LINK#",
				Value: redirectUrl,
			},
			{
				Key:   "#SUBTITLE#",
				Value: footerText,
			},
		},
	}
}

