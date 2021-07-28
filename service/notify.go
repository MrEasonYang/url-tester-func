package service

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"net/url"
	"strings"
	"log"
	"sync"
)

const TYPE_TELEGRAM = "telegram"
const TYPE_FTQQ_V1 = "ftqq_v1"
const TYPE_FTQQ_V2 = "ftqq_v2"
const TYPE_QMSG_CHAT = "qmsg_chat"
const TYPE_QMSG_GROUP_CHAT = "qmsg_group_chat"

var singleton sync.Once
var notifierMapByType = make(map[string]Notifier)

// GetNotifierByType is the way to produce concret notifier singleton instance.
func GetNotifierByType(notifierType string) Notifier {
	singleton.Do(func() {
		notifierMapByType[TYPE_TELEGRAM] = TelegramNotifier{}
		notifierMapByType[TYPE_FTQQ_V1] = FtqqV1Notifier{}
		notifierMapByType[TYPE_FTQQ_V2] = FtqqV2Notifier{}
		notifierMapByType[TYPE_QMSG_CHAT] = QmsgChatNotifier{}
		notifierMapByType[TYPE_QMSG_GROUP_CHAT] = QmsgGroupChatNotifier{}
	})	
	return notifierMapByType[notifierType]
}

// Notifier is the abstract interface for notifying users.
type Notifier interface {
	Notify(msgToSend string, config NotifierConfig)
}

// TelegramNotifier, support Telegram, see url-tester-func docs.
type TelegramNotifier struct {}

func (notifier TelegramNotifier) Notify(msgToSend string, config NotifierConfig) {
	if config.TelegramChatId == 0 || config.TelegramBot == nil {
		log.Panic("Telegram config is not valid")
	}
	msg := tgbotapi.NewMessage(config.TelegramChatId, msgToSend)
	config.TelegramBot.Send(msg)
}

// FtqqV1Notifier support WeChat .etc, see https://sct.ftqq.com/
type FtqqV1Notifier struct {}

func (notifier FtqqV1Notifier) Notify(msgToSend string, config NotifierConfig) {
	if config.FtqqV1Key == "" {
		log.Panic("FtqqV1Key is blank")
	}
	data := url.Values{}
	data.Set("text", "url-tester-func:failure")
    data.Set("desp", msgToSend)
	sendFormData("https://sc.ftqq.com/" + config.FtqqV1Key + ".send", data)
}

// FtqqV2Notifier support WeChat .etc, see https://sct.ftqq.com/
type FtqqV2Notifier struct {}

func (notifier FtqqV2Notifier) Notify(msgToSend string, config NotifierConfig) {
	if config.FtqqV1Key == "" {
		log.Panic("FtqqV2Key is blank")
	}
	data := url.Values{}
	data.Set("title", "url-tester-func:failure")
    data.Set("desp", msgToSend)
	sendFormData("https://sctapi.ftqq.com/" + config.FtqqV2Key + ".send", data)
}

// QmsgChatNotifier support QQ single chat, see https://qmsg.zendee.cn/index.html
type QmsgChatNotifier struct {}

func (notifier QmsgChatNotifier) Notify(msgToSend string, config NotifierConfig) {
	if config.QmsgKey == "" {
		log.Panic("QmsgKey is blank")
	}
	data := url.Values{}
    data.Set("msg", msgToSend)
	sendFormData("https://qmsg.zendee.cn/send/" + config.QmsgKey, data)
}

// QmsgGroupChatNotifier support QQ single chat, see https://qmsg.zendee.cn/index.html
type QmsgGroupChatNotifier struct {}

func (notifier QmsgGroupChatNotifier) Notify(msgToSend string, config NotifierConfig) {
	if config.QmsgKey == "" {
		log.Panic("QmsgKey is blank")
	}
	data := url.Values{}
    data.Set("msg", msgToSend)
	sendFormData("https://qmsg.zendee.cn/group/" + config.QmsgKey, data)
}

func sendFormData(api string, data url.Values) {
	resp, err := http.Post(
		api,
		"application/x-www-form-urlencoded", 
		strings.NewReader(data.Encode()))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
}
