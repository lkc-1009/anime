package event

import (
	TToken "anime/main/token"
	"context"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
	"time"
)

func LoadEvent() {
	botToken := token.BotToken(TToken.TConfig.AppID, TToken.TConfig.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	ws, err := api.WS(ctx, nil, "")

	if err != nil {
		log.Fatalln("websocket错误， err = ", err)
	}

	var atMessage event.ATMessageEventHandler = atMessageEventHandler
	intent := websocket.RegisterHandlers(atMessage)
	err = botgo.NewSessionManager().Start(ws, botToken, &intent)
	if err != nil {
		return
	}
}
