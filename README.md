# Go Telegram Bot


### Steps

Create a new telegram bot with sending `/newbot` command to [botfather](https://t.me/botfather)

Get chatids with `https://api.telegram.org/bot<token>/getUpdates`

Send message with `https://api.telegram.org/bot<token>/sendMessage?chat_id=<chatid>&text=<message>` to given chatid.


[TELEGRAM BOT DOCS](https://core.telegram.org/bots)

### Local Run

Create .env file and set `BOT_TOKEN` and `CHAT_ID`

Create image `docker build -t go-telegram-bot .`

Run image with .env file `docker run --env-file .env go-telegram-bot`