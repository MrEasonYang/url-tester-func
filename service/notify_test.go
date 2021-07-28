package service

import (
	"testing"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// To start testing, you have to fill these constants.
const telegram_token = ""
const telegram_chat_id int64 = 0
const ftqq_v1_key = ""
const ftqq_v2_key = ""
const qmsg_key = ""
const symbol_test_str = ""

func TestGetNotifierByType(t *testing.T) {
	notifier := GetNotifierByType(TYPE_TELEGRAM)
	_, instanceof := notifier.(TelegramNotifier)
	if !instanceof {
		t.Errorf("Not telegram notifier")
	}

	notifier = GetNotifierByType(TYPE_FTQQ_V1)
	_, instanceof = notifier.(FtqqV1Notifier)
	if !instanceof {
		t.Errorf("Not ftqqv1 notifier")
	}

	notifier = GetNotifierByType(TYPE_FTQQ_V2)
	_, instanceof = notifier.(FtqqV2Notifier)
	if !instanceof {
		t.Errorf("Not ftqqv2 notifier")
	}

	notifier = GetNotifierByType(TYPE_QMSG_CHAT)
	_, instanceof = notifier.(QmsgChatNotifier)
	if !instanceof {
		t.Errorf("Not qmsg chat notifier")
	}

	notifier = GetNotifierByType(TYPE_QMSG_GROUP_CHAT)
	_, instanceof = notifier.(QmsgGroupChatNotifier)
	if !instanceof {
		t.Errorf("Not qmsg group chat notifier")
	}
}

func TestSendToTelegram(t *testing.T) {
	notifier := GetNotifierByType(TYPE_TELEGRAM)
	bot, err := tgbotapi.NewBotAPI(telegram_token)
	if err != nil {
		t.Error(err)
	}
	notifier.Notify("Telegram msg test" + symbol_test_str, NotifierConfig{
		TelegramToken: telegram_token,
		TelegramChatId: telegram_chat_id,
		TelegramBot: bot,
	})
}

func TestSendToFtqqV1(t *testing.T) {
	notifier := GetNotifierByType(TYPE_FTQQ_V1)
	notifier.Notify("Ftqq v1 msg test" + symbol_test_str, NotifierConfig{
		FtqqV1Key: ftqq_v1_key,
	})
}

func TestSendToFtqqV2(t *testing.T) {
	notifier := GetNotifierByType(TYPE_FTQQ_V2)
	notifier.Notify("Ftqq v2 msg test" + symbol_test_str, NotifierConfig{
		FtqqV2Key: ftqq_v2_key,
	})
}

func TestSendToQmsgChat(t *testing.T) {
	notifier := GetNotifierByType(TYPE_QMSG_CHAT)
	notifier.Notify("Qmsg chat test" + symbol_test_str, NotifierConfig{
		QmsgKey: qmsg_key,
	})
}

func TestSendToQmsgGroupChat(t *testing.T) {
	notifier := GetNotifierByType(TYPE_QMSG_GROUP_CHAT)
	notifier.Notify("Qmsg group chat msg test" + symbol_test_str, NotifierConfig{
		QmsgKey: qmsg_key,
	})
}