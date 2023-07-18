module main

go 1.20

require (
	configurator v0.0.0
	initializer v0.0.0
	logs v0.0.0
	vars v0.0.0
)

require (
	github.com/my10c/packages-go/is v0.0.0-20230629045125-5efc1e3334c4 // indirect
	github.com/my10c/packages-go/lock v0.0.0-20230629045125-5efc1e3334c4
	github.com/my10c/packages-go/print v0.0.0-20230629045125-5efc1e3334c4 // indirect
	github.com/my10c/packages-go/spinner v0.0.0-20230629045125-5efc1e3334c4
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/akamensky/argparse v1.3.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/slack-go/slack v0.12.2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace configurator => ./mod/configurator

replace initializer => ./mod/initializer

replace logs => ./mod/logs

replace vars => ./mod/vars
