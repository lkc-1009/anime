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

func CreateChildChannel(guildId string, dto *dto.ChannelValueObject) *dto.Channel {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	channel, channelError := api.PostChannel(ctx, guildId, dto)
	if channelError != nil {
		log.Fatalln("调用 PostChannel 接口失败, err = ", channelError)
	}

	log.Println(channel)
	return channel
}
