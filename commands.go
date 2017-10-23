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
		EnvVar: "ENVER_VEBOSE",
		Name:   "vebose",
		Usage:  "Be verbose; show information while parsing",
	},
	cli.BoolFlag{
		EnvVar: "ENVER_DEBUG",
		Name:   "debug",
		Usage:  "Show debug logging.",
	},
	cli.BoolFlag{
		EnvVar: "ENVER_NOOP",
		Name:   "noop",
		Usage:  "Toggles a 'dry-run' only showing what would be done.",
	},
}

// Commands are all the commands supported by enver
var Commands = []cli.Command{
	{
		Name:   "parse",
		Usage:  "",
		Action: command.CmdParse,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "show",
		Usage:  "",
		Action: command.CmdShow,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound shows an error message when the user requests a command that doesn't exist
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not an %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
