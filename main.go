package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/MrEasonYang/url-tester-func/service"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func entryFunc(ctx context.Context) (string, error) {
	var targetConfigArr []service.TargetConfig
    config := os.Getenv("config")
    if config == "" {
        return "Config is blank", errors.New("Invalid param")
    }
	err := json.Unmarshal([]byte(config), &targetConfigArr)
    if err != nil {
        return "Failed to parse config str to json", err
    }

    rawChatIDConfig := os.Getenv("chatID")
    if rawChatIDConfig == "" {
        return "ChatID is blank", errors.New("Invalid param")
    }
    chatID, err := strconv.ParseInt(rawChatIDConfig, 10, 64)
    if err != nil {
        return "Failed to parse chatID", err
    }

    telegramToken := os.Getenv("token")
    if telegramToken == "" {
        return "TelegramToken is blank", errors.New("Invalid param")
    }

	return service.TesterHandler(targetConfigArr, telegramToken, chatID), nil
}

func main() {
	cloudfunction.Start(entryFunc)
}
