package message

import (
	"github.com/tencent-connect/botgo/dto"
	"strings"
)

func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	if strings.Compare(data.Content, "> hello") == 0 {
		_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{Ark: createEmbed()})
		if err != nil {
			return err
		}
	}
	return nil
}
