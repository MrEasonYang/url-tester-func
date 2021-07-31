package service

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// TargetConfig is the info pair of the target url to be tested.
type TargetConfig struct {
	URL                string `json:"url"`
	ExpectedStatusCode int `json:"expectedStatusCode"`
	IgnoreAnalysis     bool `json:"ignoreAnalysis"`
	NotifyMethod       string `json:"notifyMethod"`
}

type NotifierConfig struct {
	FtqqV1Key            string
	FtqqV2Key            string
	QmsgKey              string
	TelegramToken        string
	TelegramChatId       int64
	TelegramBot          *tgbotapi.BotAPI
}

// Result is the return msg combination struct for the tester.
type Result struct {
	URL     string
	Msg     string
	Succeed bool
}
