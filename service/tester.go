package service

import (
	"net/http"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"fmt"
	"encoding/json"
)

var bot *tgbotapi.BotAPI
var botChatID int64

// TesterHandler is the entry func of the application.
func TesterHandler(targetArr []TargetConfig, telegramToken string, telegramChatID int64) string {
	var err error
	bot, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	botChatID = telegramChatID
	byteData, _ := json.Marshal(doRequestTest(targetArr))
	return string(byteData)
}

// Asynchronously request all target and report errors.
func doRequestTest(targetArr []TargetConfig) []Result {
	if len(targetArr) == 0 {
		return nil
	}

	httpClient := &http.Client{}

	ch := make(chan Result, len(targetArr))
	counter := make(chan int, len(targetArr))

	for _, targetConfig := range targetArr {
		currentTargetConfig := targetConfig

		go func() {
			proxyReq, err := http.NewRequest(http.MethodGet, currentTargetConfig.URL, nil)

			proxyReq.Header = make(http.Header)
			proxyReq.Header.Set("user-agent", "url-tester-func")

			resp, err := httpClient.Do(proxyReq)
			if err != nil {
				msg := fmt.Sprintf("Failed to test URL %s due to error %s", currentTargetConfig.URL, err)
				sendAlertMsg(msg)
				ch <- Result{URL: currentTargetConfig.URL, Msg: msg, Succeed: false}
				counter <- 1
				return
			}
			if resp.StatusCode != currentTargetConfig.ExpectedStatusCode {
				msg := fmt.Sprintf("Failed to test URL %s due to status code is %d rather than %d", 
					currentTargetConfig.URL, resp.StatusCode, currentTargetConfig.ExpectedStatusCode)
				sendAlertMsg(msg)
				ch <- Result{URL: currentTargetConfig.URL, Msg: msg, Succeed: false}
				counter <- 1
				return
			}
			ch <- Result{URL: currentTargetConfig.URL, Msg: "", Succeed: true}
			counter <- 1
			defer resp.Body.Close()
		}()
	}

	var result []Result
	go func ()  {
		sum := 0
		for v := range counter {
			sum += v
			if sum == len(targetArr) {
				close(ch)
				close(counter)
			}
		}
	}()
	for v := range ch {
		result = append(result, v)
	}
	return result
}

// Send alert via telegram.
func sendAlertMsg(msgToSend string) {
	msg := tgbotapi.NewMessage(botChatID, msgToSend)
	bot.Send(msg)
}