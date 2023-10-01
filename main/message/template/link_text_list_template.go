package template

import "github.com/tencent-connect/botgo/dto"

// CreateLinkTextListTemplate 创建链接+文本列表模板
// obj -> {{Key: "desc/link", Value: "text/linkText,}, {}...}
func createLinkTextListTemplate(
	desc string,
	tips string,
	obj []*dto.ArkObjKV,
) *dto.Ark {
	return &dto.Ark{
		TemplateID: 23,
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
				Key: "#LIST#",
				Obj: []*dto.ArkObj{
					{
						obj,
					},
				},
			},
		},
	}
}
