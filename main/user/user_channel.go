package user

import (
	TToken "anime/main/token"
	"context"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/token"
	"log"
	"time"
)

func UserChannelInfo() []*dto.Guild {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	guilds, meGuildError := api.MeGuilds(ctx, &dto.GuildPager{})
	if meGuildError != nil {
		log.Fatalln("调用 MeGuild 接口失败, err = ", meGuildError)
	}

	log.Println(guilds)
	return guilds
}
