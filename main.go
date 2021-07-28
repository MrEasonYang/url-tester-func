package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"github.com/MrEasonYang/url-tester-func/service"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func entryFunc(ctx context.Context) (string, error) {
	var targetConfigArr []service.TargetConfig
	config := os.Getenv("config")
	if config == "" {
		return "Config is blank", errors.New("invalid param")
	}
	err := json.Unmarshal([]byte(config), &targetConfigArr)
	if err != nil {
		return "Failed to parse config str to json", err
	}

	var telegramChatID int64 = 0
	rawChatIDConfig := os.Getenv("telegram_chat_id")
	if rawChatIDConfig != "" {
		var parseErr error
		telegramChatID, parseErr = strconv.ParseInt(rawChatIDConfig, 10, 64)
		if parseErr != nil {
			return "Failed to parse chatID", err
		}
	}

	telegramToken := os.Getenv("telegram_token")
	ftqqV1Key := os.Getenv("ftqq_v1_key")
	ftqqV2Key := os.Getenv("ftqq_v2_Key")
	qmsgKey := os.Getenv("qmsg_key")

	notifierConfig := service.NotifierConfig {
		TelegramChatId: telegramChatID,
		TelegramToken: telegramToken,
		FtqqV1Key: ftqqV1Key,
		FtqqV2Key: ftqqV2Key,
		QmsgKey: qmsgKey,
	}

	if telegramToken != "" && telegramChatID != 0 {
		bot, err := tgbotapi.NewBotAPI(telegramToken)
		if err != nil {
			return "Failed to init telegram bot client", err
		}
		notifierConfig.TelegramBot = bot
	}

	return service.TesterHandler(targetConfigArr, notifierConfig), nil
}

func main() {
	cloudfunction.Start(entryFunc)
}
