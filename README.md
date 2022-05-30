# url-tester-func
A simple url tester function based on Golang and Serverless.
[中文说明](https://easonyang.com/2021/02/21/golang-qcloud-serverless-url-tester-func/)

# Features
- Time based url tester which is able forward the test result to the specified IM.
- Support to nofiry failure via **Telegram/Ftqq(Wechat)/Qmsg(QQ)**.
- Use specified User-Agent contains Bot description(url-tester-func BOT).
- Support QCloud Serverless(SCF).

# Usage
## Setup config
Setup a QCloud serverless config, see the example here: [serverless.yaml](https://github.com/MrEasonYang/url-tester-func/blob/main/serverless.yaml.example)

## Build
Build the project with golang:

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

## Local tests
See [QCloud SCF Readme](./README-QCLOUD_EN.md)([腾讯云 SCF 中文说明](./README-QCLOUD.md))

## Create a new function
### Add function
First of all, you have to register a QCloud account.

Then visit [function management panel](https://console.cloud.tencent.com/scf/list-create?rid=5&ns=default&createType=empty), create a new function.

**Important, to request Telegram successfully, you have to pick up a region outside China mainland.**

Then select Custom Create and here is a basic configuration example:
- Function type: Event function
- Region: Hong Kong
- Deploy: By code
- Runtime: Go1
- Memory: 128MB
- Timeout: 60 seconds
- Environment: will be explained below
- Public Network: true

And leave the other things as default.

### Setup Serverless environment
Set the global environment on QCloud platform interface:
- **config(Required)**: 
Config structure
    1. **url**: The target url to test with.
    2. **expectedStatusCode**: The HTTP response status code which should not trigger a failure notify.
    3. **notifyMethod**: The IM to send notify message, should be one of telegram/ftqq_v1/ftqq_v2/qmsg_chat/qmsg_group_chat.

Example:
```json
[
    {
        "url": "<Target URL 1 with Telegram Notifier>",
        "expectedStatusCode": 200,
        "notifyMethod": "telegram"
    },
    {
        "url": "<Target URL 2 with FtqqV1 Notifier>",
        "expectedStatusCode": 403,
        "notifyMethod": "ftqq_v1"
    },
    {
        "url": "<Target URL 3 with FtqqV2 Notifier>",
        "expectedStatusCode": 400,
        "notifyMethod": "ftqq_v2"
    },
    {
        "url": "<Target URL 4 with Qmsg Notifier>",
        "expectedStatusCode": 301,
        "notifyMethod": "qmsg_chat"
    },
    {
        "url": "<Target URL 4 with Qmsg Group Notifier>",
        "expectedStatusCode": 200,
        "notifyMethod": "qmsg_group_chat"
    },
]
```
- **telegram_token(Required when using telegram)**: Telegram bot api token, could be obtained from BotFather. Check this [FAQ](https://telegra.ph/Awesome-Telegram-Bot-11-11) out.
- **telegram_chat_id(Required when using telegram)**: Target chat associated to the telegram bot, could be obtained from the updates api, see telegram docs for more information. Here is a tutorial: [How to get a group chat id?](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)
- **qmsg_key(Required when using qmsg)**: [qmsg key](https://qmsg.zendee.cn/me.html)
- **ftqq_v1_key(Required when using ftqq v1 api)**: [ftqq send key](https://sct.ftqq.com/sendkey), only users registered before V2 released can obtain the send key due to the V1 API is deprecated now by ftqq.
- **ftqq_v2_Key(Required when using ftqq v2 api)**: [ftqq send key](https://sct.ftqq.com/sendkey)

### Upload codes
Use ZIP(recommended) to upload the entire function directory including the executable file we've just built.

### Add a time based trigger
Setup a time based trigger, for example, using these parameters if you wish to run the tester every 5 minutes: 
- Trigger period: Every 5 minutes
- Cron expression: `0 */5 * * * * *`

**Then run it and enjoy it.**

## Todo
Currently url-tester-func can only work under QCloud SCF, more platforms should be supported in the future:
- Aliyun Serverless
- AWS Lambda
- Azure Serverless
- GCP Serverless

## About the bills
<del>It will be free for small usage like less then 10 urls and 10k invocations per month, however, **the more resources you use the more money you will have to pay** because the outcome public internet requests are always not free. See QCloud official [docs](https://cloud.tencent.com/product/scf/pricing) for more information.</del>

**Warning: Using Tencent cloud serverless will be charged with a minimum cost since 2022/06/01.**

## Contribution
Contributions are welcome, just remember to lint your code.

## Credits
[go-telegram-bot-api/telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)

## License
[MIT](https://github.com/MrEasonYang/url-tester-func/blob/main/LICENSE)
