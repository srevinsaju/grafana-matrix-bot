# grafana-matrix-bot 🤖
Grafana Matrix Bot to send webhook notifications


## Usage 🤔

```bash
$ ./grafana-matrix-bot path/to/config.json
```

## Installation ✨

1. Download the precompiled binary from [Releases](https://github.com/srevinsaju/grafana-matrix-bot/releases) or
   alternatively, [build from source](#build).
2. Copy `config.sample.json` to `config.json` and edit the values appropriately
3. Start the bot by
```bash
./grafana-matrix-bot config.json
```


## Build 🔧

```bash
git clone https://github.com/srevinsaju/grafana-matrix-bot.git && cd grafana-matrix-bot
go build .
```


## License

See [LICENSE](./LICENSE) for more information.
