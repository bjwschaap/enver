package command

import (
	"fmt"
	"log"

	"github.com/bjwschaap/enver/config"
	"github.com/urfave/cli"
)

// CmdShow implements the 'show' command
func CmdShow(c *cli.Context) error {
	config, err := config.FromFile(c)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	files, err := config.FilesToParse(c)
	if err != nil {
		if c.GlobalBool("debug") || config.Debug {
			log.Printf("error while making list of files to parse: %v", err)
		}
	}

	log.Printf("Enver config file: %s", c.GlobalString("config"))
	log.Printf("Debug mode: %t", c.GlobalBool("debug") || config.Debug)
	log.Printf("Quiet: %t", c.GlobalBool("quiet") || config.Quiet)
	log.Println("Files:")
	for _, f := range files {
		log.Printf( "\t- %s", f)
	}

	return nil
}
