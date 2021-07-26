# url-tester-func
A simple url tester function based on Golang and Serverless.
[中文说明](https://easonyang.com/2021/02/21/golang-qcloud-serverless-url-tester-func/)

# Features
- Time based url tester which is able forward the test result to the specified telegram chat.
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
- **config**: 
```json
[{"URL": "<Target URL 1>","ExpectedStatusCode": 200},{"URL": "<Target URL 2>","ExpectedStatusCode": 403}]
```
- **token**: Telegram bot api token, could be obtained from BotFather. Check this [FAQ](https://telegra.ph/Awesome-Telegram-Bot-11-11) out.
- **chatID**: Target chat associated to the telegram bot, could be obtained from the updates api, see telegram docs for more information. Here is a tutorial: [How to get a group chat id?](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)

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

## Contribution
Contributions are welcome, just remember to lint your code.

## License
[MIT](https://github.com/MrEasonYang/url-tester-func/blob/main/LICENSE)