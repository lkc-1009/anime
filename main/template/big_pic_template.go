package template

import "github.com/tencent-connect/botgo/dto"

// CreateBigPicTemplate 创建大图模板
func CreateBigPicTemplate(
	tips string,
	title string,
	subTitle string,
	picUrl string,
	redirectUrl string,
) *dto.Ark {
	return &dto.Ark{
		TemplateID: 37,
		KV: []*dto.ArkKV{
			{
				Key:   "#PROMPT#",
				Value: tips,
			},
			{
				Key:   "#METATITLE#",
				Value: title,
			},
			{
				Key:   "#METASUBTITLE#",
				Value: subTitle,
			},
			{
				Key:   "#METACOVER#",
				Value: picUrl,
			},
			{
				Key:   "#METAURL#",
				Value: redirectUrl,
			},
		},
	}
}
