# Just Another Meme Sticker (JAMeS)

A friend sent you a meme in Telegram, and it is actually funny. You want to turn that piece of contemporary art into a sticker using the default way is through the @stickers bot, but then:
- "*Please check that the image fits into a 512x512 square (one of the sides should be 512px and the other 512px or less).*"
- "*Please attach the image as a file (uncompressed), not as a photo.*"
- "*Sorry, the file type is invalid. Please convert your image to PNG.*" (the file format of photo in chat)

The requirements are quite troublesome, especially for mobile users that needs to perform these steps manually that could take several minutes. And thus, this bot is created as an attempt (successful one) to automate the process.
This bot is made using [Go](https://go.dev) and [Firebase](https://firebase.google.com).

## Setup
### Firebase
If you are still planning to use Firebase as the database, you need to optain the service account key (Project Settings > Service Accounts > Generate new private key) and put it in the project directory. Then, in the environment file, you need to put the path to the service account key in the `FIREBASE_SERVICEACCOUNT_KEY_PATH` variable.

### Telegram
You need to create a bot using [BotFather](https://t.me/botfather) and put the bot token in the `TELEGRAM_APITOKEN` variable.

## Installation

Once the setup is done, you can start by installing Go from [here](https://go.dev/doc/install) (tested in version `1.20.4`).

If it is installed, you can run install the required packages and build an executable file by running:
```bash
go install
go build
```

Once it is done installing the dependencies, you can build the bot by running:
```bash
./james
```

## Commands
- `/start` - Start the bot
- `/initialize` - Initialize the sticker pack in the chat (one sticker pack per chat only)
- `/add` - Add a sticker to the sticker pack
- `/getpack` - Get the sticker pack link for the current chat
- `/connect` - Connect an existing sticker pack (made by through this bot) to the chat
- `/disconnect` - Disconnect the sticker pack from the chat
- `/help` - Show the help message

To check the parameters of the sticker pack, you can just type the command without any parameters.

# Suggestions
If you have any suggestions or input, feel free to contact me through [my email](mailto:claytonfernalo@gmail.com).&nbsp;

<img src="https://media.tenor.com/o-0LaJK3qWcAAAAC/yamada-ryou-yamada-ryo.gif" width="360">

# Changelogs
- `(8/1/2023)` First version


