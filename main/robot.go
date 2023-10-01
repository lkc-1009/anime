package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AppID uint64 `yaml:"appid"`
	Token string `yaml:"token"`
}

var config Config
var api openapi.OpenAPI
var ctx context.Context

func init() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("读取配置文件出错， err = ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Println("解析配置文件出错， err = ", err)
		os.Exit(1)
	}
	log.Println(config)
}

func main() {
	robotToken := token.BotToken(config.AppID, config.Token)
	api = botgo.NewOpenAPI(robotToken).WithTimeout(3 * time.Second)
	ctx = context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln("websocket错误， err = ", err)
	}

	var atMessage event.ATMessageEventHandler = atMessageEventHandler

	intent := websocket.RegisterHandlers(atMessage)
	err = botgo.NewSessionManager().Start(ws, robotToken, &intent)
	if err != nil {
		return
	}
}

func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	if strings.HasSuffix(data.Content, "> hello") {
		_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{Ark: createEmbed()})
		if err != nil {
			return err
		}
	}
	return nil
}

func createEmbed() *dto.Ark {
	return &dto.Ark {
		TemplateID: 37,
		KV: []*dto.ArkKV {
			{
				Key: "#PROMPT#",
				Value: "你好",
			},
			{
				Key: "#METATITLE#",
				Value: "你好",
			},
			{
				Key: "#METASUBTITLE#",
				Value: "你好",
			},
		},
	}
}
