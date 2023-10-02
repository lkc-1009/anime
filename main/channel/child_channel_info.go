package channel

import (
	TToken "anime/main/token"
	"context"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/token"
	"log"
	"time"
)

func ChildChannelInfo(channelID string) *dto.Channel {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	channel, channelError := api.Channel(ctx, channelID)
	if channelError != nil {
		log.Fatalln("调用 Channel 接口失败, err = ", channelError)
	}

	log.Println(channel)
	return channel
}
