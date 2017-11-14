package main

import (
	"fmt"
	"os"

	"github.com/bjwschaap/enver/command"
	"github.com/urfave/cli"
)

// GlobalFlags are the global flags that can be set for the program
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "ENVER_QUIET",
		Name:   "quiet, q",
		Usage:  "Don't show any output.",
	},
	cli.BoolFlag{
		EnvVar: "ENVER_DEBUG",
		Name:   "debug, d",
		Usage:  "Show debug logging.",
	},
	cli.StringFlag{
		EnvVar: "ENVER_CONFIG",
		Name:   "config, c",
		Usage:  "Load configuration from `FILE`",
		Value:  ".enver",
	},
}

// Commands are all the commands supported by enver
var Commands = []cli.Command{
	{
		Name:   "parse",
		Usage:  "Parse the files and replace placeholders with values from ENV",
		Action: command.CmdParse,
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				EnvVar: "ENVER_FILES",
				Name:   "files, f",
				Usage:  "A list of files that enver should parse.",
			},
			cli.BoolFlag{
				EnvVar: "ENVER_NOOP",
				Name:   "noop",
				Usage:  "Toggles a 'dry-run' only showing what would be done.",
			},
		},
	},
	{
		Name:   "show",
		Usage:  "Show the enver configuration",
		Action: command.CmdShow,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound shows an error message when the user requests a command that doesn't exist
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not an %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
