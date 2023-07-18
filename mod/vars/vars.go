//
// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package vars

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/my10c/packages-go/is"
	"github.com/my10c/packages-go/print"
)

var (
	Off    = "\x1b[0m"    // Text Reset
	Black  = "\x1b[1;30m" // Black
	Red    = "\x1b[1;31m" // Red
	Green  = "\x1b[1;32m" // Green
	Yellow = "\x1b[1;33m" // Yellow
	Blue   = "\x1b[1;34m" // Blue
	Purple = "\x1b[1;35m" // Purple
	Cyan   = "\x1b[1;36m" // Cyan
	White  = "\x1b[1;37m" // White

	RedUnderline = "\x1b[4;31m" // Red underline
	OneLineUP    = "\x1b[A"
)

type Funcs struct {
	I *is.Is
	P *print.Print
}

// for logging
type Log struct {
	LogsDir       string
	LogFile       string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
}

// Configuration for Slack
type SlackConfig struct {
	SlackToken 	[]string
	SlackUser   []string
	SlackMsg    []string
	SlackEmoji  []string
}

var (
	MyVersion   = "0.0.1"
	now         = time.Now()
	MyProgname  = path.Base(os.Args[0])
	myAuthor    = "Luc Suryo"
	myCopyright = "Copyright 2023 - " + strconv.Itoa(now.Year()) + " ©Badassops LLC"
	myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
	myEmail     = "<luc@badassops.com>"
	MyInfo      = fmt.Sprintf("%s (version %s)\n%s\n%s\nWritten by %s %s\n",
		MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)
	MyDescription = "Simple script send a message to a slack channel"

	// Default configuration file
	ConfigFile = "/usr/local/etc/send-to-slack/config.ini"

	// default values for logs
	LogEnable	bool = false
	Logs		Log

	// default emoji
	Emoji = "🚨"

	// default value for lock
	Lock		bool = false
	LockFile	= "/tmp/" + MyProgname + ".lock"

	// default quiet
	Quiet		bool = false

	// we sets these under variable
	// default values
	LogsDir       = fmt.Sprintf("/var/log/%s", MyProgname)
	LogFile       = fmt.Sprintf("%s.log", MyProgname)
	LogMaxSize    = 128 // megabytes
	LogMaxBackups = 14  // 14 files
	LogMaxAge     = 14  // 14 days
)
