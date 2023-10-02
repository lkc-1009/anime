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

func UserInfo() *dto.User {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()
	user, meError := api.Me(ctx)
	if meError != nil {
		log.Fatalln("调用 Me 接口失败, err = ", meError)
	}

	log.Println(user)
	return user
}
