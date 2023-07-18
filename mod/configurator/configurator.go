//
// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package configurator

import (
	"fmt"
	"os"

	"vars"

	"github.com/my10c/packages-go/is"
	"github.com/my10c/packages-go/print"

	"github.com/BurntSushi/toml"
	"github.com/akamensky/argparse"
)

type (
	Config struct {
		AuthValues		Auth
		SlackValues		SlackConfig
		LogValues		LogConfig
		// given from the command line
		SlackMessage	[]string
		// from command line or default is used
		ConfigFile		string
		MsgEmoji		string
		LockFile		string
		LockPID			int
		Quite			bool
	}

	LogConfig struct {
		LogEnable		bool
		LogsDir			string
		LogFile			string
		LogMaxSize		int
		LogMaxBackups	int
		LogMaxAge		int
	}

	SlackConfig struct {
		Token		string
		User		string
		Channel		string
		UserEmoji	string
		MsgEmoji	string
		Lock		bool
		LockFile	string
	}

	Auth struct {
		AllowUsers	[]string
		AllowMods	[]string
	}

	tomlConfig struct {
		Auth		Auth			`toml:"auth"`
		Slack		SlackConfig		`toml:"slack"`
		LogConfig	LogConfig		`toml:"logconfig"`
	}
)

var (
	Is		= is.New()
	Print = print.New()
)

// function to initialize the configuration
func Configurator() *Config {
	// the rest of the values will be filled from the given configuration file
	return &Config{}
}

func (c *Config) InitializeArgs() {
	parser := argparse.NewParser(vars.MyProgname, vars.MyDescription)
	configFile := parser.String("c", "configFile",
		&argparse.Options{
			Required: false,
			Help:		"Configuration file to be use",
			Default:	vars.ConfigFile,
		})

	slackMessage := parser.StringList("m", "message",
		&argparse.Options{
			Required: false,
			Help:		"Message to be sent between double quotes or single quotes, required",
		})

	slackEmoji := parser.String("e", "emoji",
		&argparse.Options{
			Required:	false,
			Help:		"Emoji to use.",
		})

	quietFlag := parser.Flag("q", "quiet",
		&argparse.Options{
			Required:	false,
			Help:		"quiet mode",
			Default:	vars.Quiet,
		})

	showVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:		"Show version",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *showVersion {
		Print.ClearScreen()
		Print.PrintYellow(vars.MyProgname + " version: " + vars.MyVersion + "\n")
		os.Exit(0)
	}

	if len(*slackMessage) == 0 {
		Print.PrintRed("The flag -m/--message is required\n")
		os.Exit(1)
	} 

	if _, ok, _ := Is.IsExist(*configFile, "file"); !ok {
		Print.PrintRed("Configuration file " + *configFile + " does not exist\n")
		os.Exit(1)
	}

	if *quietFlag {
		c.Quite = true
	}

	c.ConfigFile = *configFile
	c.SlackMessage = *slackMessage
	c.MsgEmoji = *slackEmoji
}

// function to add the values to the Config object from the configuration file
func (c *Config) InitializeConfigs() {
	var configValues tomlConfig
	
	// set default value and then overwrite if exist in the configuration file
	// set to default for lock
	c.SlackValues.Lock = vars.Lock
	c.SlackValues.LockFile = vars.LockFile

	// set to default for log
	c.LogValues.LogEnable = vars.LogEnable
	c.LogValues.LogFile = vars.LogFile

	if _, err := toml.DecodeFile(c.ConfigFile, &configValues); err != nil {
		Print.PrintRed("Error reading the configuration file\n")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(configValues.Slack.Token) == 0 ||
		len(configValues.Slack.User) == 0 ||
		len(configValues.Slack.Channel) == 0 {
		Print.PrintRed("Error reading the configuration file, some value are missing or is empty\n")
		Print.PrintBlue("Make sure token, user and channel are set\n")
		Print.PrintBlue("Aborting...\n")
		os.Exit(1)
	}

	c.AuthValues = configValues.Auth
	c.LogValues = configValues.LogConfig
	c.SlackValues = configValues.Slack
	c.LockFile = vars.LockFile
}
