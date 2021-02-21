# url-tester-func
A simple url tester function based on qcloud serverless.

## Usage
1. Setup a Qcloud serverless config, see the example here: [serverless.yaml](https://github.com/MrEasonYang/url-tester-func/blob/main/serverless.yaml.example)

2. Build the project:

   ```shell
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
   ```

3. Set the global environment on Qcloud platform interface:

   - **config**: [{"URL": "<Target URL 1>","ExpectedStatusCode": 200},{"URL": "<Target URL 2>","ExpectedStatusCode": 403}]
   - **token**: <Telegram bot api token, could be obtained from BotFather>
   - **chatID**: <Target chat associated to the telegram bot, could be obtained from the updates api, see telegram docs for more information>

4. Upload to Qcloud platform, setup a time based trigger then run it and enjoy it.

## License

[MIT](https://github.com/MrEasonYang/url-tester-func/blob/main/LICENSE)