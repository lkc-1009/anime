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

func ChildChannelList(guildId string) []*dto.Channel {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	channels, channelsError := api.Channels(ctx, guildId)
	if channelsError != nil {
		log.Fatalln("调用 Channels, err = ", channelsError)
	}

	log.Println(channels)
	return channels
}
