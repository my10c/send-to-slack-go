// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	// local
	"configurator"
	"initializer"
	"logs"
	"vars"

	// on github
	"github.com/my10c/packages-go/lock"
	"github.com/my10c/packages-go/spinner"
	"github.com/slack-go/slack"
)

func main() {
	LockPid := os.Getpid()
	progName, _ := os.Executable()
	progBase := filepath.Base(progName)

	s := spinner.New(10)

	config := configurator.Configurator()

	// get given parameters
	config.InitializeArgs()

	// get the configuration
	config.InitializeConfigs()

	// initialize the value to default if not defined
	funcs := initializer.Init(config)

	// make sure the configuration file has the proper settings
	runningUser, _ := funcs.I.IsRunningUser()
	if !funcs.I.IsInList(config.AuthValues.AllowUsers, runningUser) {
		funcs.P.PrintRed(fmt.Sprintf("The program has to be run as these user(s): %s or use sudo, aborting..\n",
		strings.Join(config.AuthValues.AllowUsers[:], ", ")))
		os.Exit(0)
	}
	ownerInfo, ownerOK := funcs.I.IsFileOwner(config.ConfigFile, config.AuthValues.AllowUsers)
	if !ownerOK {
		funcs.P.PrintRed(fmt.Sprintf("%s,\nAborting..\n", ownerInfo))
		os.Exit(0)
	}
	permInfo, permOK := funcs.I.IsFilePermission(config.ConfigFile, config.AuthValues.AllowMods)
	if !permOK {
		funcs.P.PrintRed(fmt.Sprintf("%s,\nAborting..\n", permInfo))
		os.Exit(0)
	}

	if !config.Quite {
		go s.Run()
	}

	// initialize the logger system is it was set to true
	if config.LogValues.LogEnable {
		LogConfig := &logs.LogConfig{
			LogsDir:		config.LogValues.LogsDir,
			LogFile:		config.LogValues.LogFile,
			LogMaxSize:		config.LogValues.LogMaxSize,
			LogMaxBackups:	config.LogValues.LogMaxBackups,
			LogMaxAge:		config.LogValues.LogMaxAge,
		}

		logs.InitLogs(LogConfig)
		logs.Log("System all clear", "INFO")
	}

	// prevent a race
	time.Sleep(1 * time.Second)
	if !config.Quite {
		s.Stop()
	}

	if config.SlackValues.Lock {
		// create the lock file to prevent an other script is running/started if lock was set
		lockPtr := lock.New(config.LockFile)
		// check lock file; lock file should not exist
		config.LockPID = LockPid
		if _, fileExist, _ := funcs.I.IsExist(config.LockFile, "file"); fileExist {
	 		lockPid, _ := lockPtr.LockGetPid()
			if progRunning, _ := funcs.I.IsRunning(progBase, lockPid); progRunning {
	 			funcs.P.PrintRed(fmt.Sprintf("\nError there is already a process %s running, aborting...\n", progBase))
				os.Exit(0)
			}
		}
		// save to create new or overwrite the lock file
		if err := lockPtr.LockIt(LockPid); err != nil {
			funcs.P.PrintRed(fmt.Sprintf("\nError creating the lock file, error %s, aborting..\n", err.Error()))
			os.Exit(0)
		}
	}

	// prepare the message
	slackMsg := fmt.Sprintf("%s %s\n",config.SlackValues.MsgEmoji, strings.Join(config.SlackMessage, " "))
	// create the slack object
	slackAPI := slack.New(config.SlackValues.Token)
	// setup the message options
	slackMsgOptions := slack.PostMessageParameters{
		Username:       config.SlackValues.User,
        IconEmoji:      config.SlackValues.UserEmoji,
        Markdown:       false,
        EscapeText:     false,
	}

	// send the message
	_, _, err := slackAPI.PostMessage(config.SlackValues.Channel,
					slack.MsgOptionText(slackMsg, false),
					slack.MsgOptionPostMessageParameters(slackMsgOptions),)
	if err !=nil {
			funcs.P.PrintRed(fmt.Sprintf("\nError sending the message, error %s..\n", err.Error()))
	}

	if !config.Quite {
		funcs.P.TheEnd()
		fmt.Printf("\t%s\n", funcs.P.PrintLine(vars.Purple, 50))
	}

	if config.SlackValues.Lock {
		os.Remove(config.LockFile)
	}

	if config.LogValues.LogEnable {
		logs.Log("System Normal shutdown", "INFO")
	}
	os.Exit(0)
}
