# send-to-slack-go
send message to a configured slack channel

## usage

```
usage: send-to-slack [-h|--help] [-c|--configFile "<value>"] [-m|--message
                     "<value>" [-m|--message "<value>" ...]] [-e|--emoji
                     "<value>"] [-q|--quiet] [-v|--version]

                     Simple script send a message to a slack channel

Arguments:

  -h  --help        Print help information
  -c  --configFile  Configuration file to be use. Default:
                    /usr/local/etc/send-to-slack/config.ini
  -m  --message     Message to be sent between double quotes or single quotes,
                    required
  -e  --emoji       Emoji to use.. Default: 🚨
  -q  --quiet       quiet mode. Default: false
  -v  --version     Show version
```

# how to build

```
go mod init
go mod tidy
go build -o send-to-slack send-to-slack.go
```
