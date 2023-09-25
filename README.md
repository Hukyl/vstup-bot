# vstup-bot
A Telegram bot for Vstup NaUKMA entry test. It is written in Golang.
**Note**: every input for currency is in [ISO-4217 format](https://en.wikipedia.org/wiki/ISO_4217).

List of commands:
- /exists \`iso\` - checks whether the currency is supported
- /list - outputs a list of the most popular (not all) supported currencies
- /convert \`iso\` \`amount\` - convert some amount of UAH to required currency 

Each command has validation checks to ensure input data correctness. 

Even though the task did not require conversion from any currency other than UAH, the bot is easily adjustable to do so.

## Running

The executable requires .env file with set `VSTUP_TOKEN` Telegram bot token.

#### Building from source
1. Download the go 1.21.1 version, the method is noted in [documentation](https://go.dev/doc/install).
2. Clone the repo.
```~# git clone https://github.com/Hukyl/vstup-bot.git```
3. Get all required libraries, mentioned in the `go.mod`.
```~/vstup-bot# go mod vendor```
4. Build and run the executable
```~/vstup-bot# go build -o ./vstup-bot && ./vstup-bot```


#### Downloading from GitHub Actions
1. Go to Actions section
2. Click latest workflow run
3. Download Linux binary (and test results, optionally)
4. Run the executable