package main

import (
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "bjwschaap"
	app.Email = "bastiaan.schaap@gmail.com"
	app.Usage = `Enver simply parses files (specified in enver.conf), and replaces
	placeholders in the form: ${env.MY_VAR} with the value found in environment
	variable MY_VAR. Optionally a default value can be provided using the pipe
	symbol. E.g.: ${env.MY_VAR|1234}. If there is no environment variable MY_VAR
	enver will replace with default value '1234'. Make enver very quiet by passing
	-q.`

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "Only print version, and exit.",
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
