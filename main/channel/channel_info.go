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

func ChannelInfo(guildId string) *dto.Guild {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	guild, guildError := api.Guild(ctx, guildId)
	if guildError != nil {
		log.Fatalln("调用 Guild 接口失败, err = ", guildError)
	}

	log.Println(guild)
	return guild
}
