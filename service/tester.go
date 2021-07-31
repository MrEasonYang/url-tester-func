package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TesterHandler is the entry func of the application.
func TesterHandler(targetArr []TargetConfig, notifierConfig NotifierConfig) string {
	byteData, _ := json.Marshal(doRequestTest(targetArr, notifierConfig))
	return string(byteData)
}

// Asynchronously request all target and report errors.
func doRequestTest(targetArr []TargetConfig, notifierConfig NotifierConfig) []Result {
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
			if err != nil {
				log.Panic(err)
			}

			proxyReq.Header = make(http.Header)
			proxyReq.Header.Set("user-agent", "url-tester-func")
			if currentTargetConfig.IgnoreAnalysis {
				proxyReq.Header.Set("analysis-action", "ignore")
			}

			resp, err := httpClient.Do(proxyReq)
			if err != nil {
				msg := fmt.Sprintf("Failed to test URL %s due to error %s", currentTargetConfig.URL, err)
				sendAlertMsg(msg, currentTargetConfig.NotifyMethod, notifierConfig)
				ch <- Result{URL: currentTargetConfig.URL, Msg: msg, Succeed: false}
				counter <- 1
				return
			}
			if resp.StatusCode != currentTargetConfig.ExpectedStatusCode {
				msg := fmt.Sprintf("Failed to test URL %s due to status code is %d rather than %d",
					currentTargetConfig.URL, resp.StatusCode, currentTargetConfig.ExpectedStatusCode)
				sendAlertMsg(msg, currentTargetConfig.NotifyMethod, notifierConfig)
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
	go func() {
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

// Send alert via specified notifier.
func sendAlertMsg(msgToSend string, notifierType string, notifierConfig NotifierConfig) {
	notifier := GetNotifierByType(notifierType)
	notifier.Notify(msgToSend, notifierConfig)
}
