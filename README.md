# Just Another Meme Sticker (JAMeS)
## Setup
### Firebase
If you are still planning to use Firebase as the database, you need to obtain the service account key (Project Settings > Service Accounts > Generate new private key) and put it in the project directory. Then, in the environment file, you need to put the path to the service account key in the `FIREBASE_SERVICEACCOUNT_KEY_PATH` environment variable.

### Telegram
You need to create a bot using [BotFather](https://t.me/botfather) and put the bot token in the `TELEGRAM_APITOKEN` environment variable.

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

## TODO:
- Concurrency (especially on image processing)
  
## Changelogs and Commands
Can be found [here](https://sanstzu.vercel.app/blogs/james-telegram-bot).


