package command

import (
	"io/ioutil"
	"fmt"
	"log"
	"regexp"
	"bytes"
	"os"

	"github.com/bjwschaap/enver/config"
	"github.com/urfave/cli"
)

// CmdParse implements the enver 'parse' command
func CmdParse(c *cli.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("current dir: %s", dir)

	// Match any placeholder with form {{env.VAR}} and/or {{env.VAR|value}}
	r := regexp.MustCompile(`{{env\.([\w\-]*)\|?(.*)?}}`)

	config, err := config.FromFile(c)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	files, err := config.FilesToParse(c)
	if err != nil {
		return fmt.Errorf("error while making list of files to parse: %v", err)
	}

	if c.GlobalBool("debug") || config.Debug {
		log.Printf("files to parse: %v", files)
	}

	for _, file := range files {
		if !c.GlobalBool("quiet") || config.Quiet {
			log.Printf("parsing %s", file)
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}

		m := r.FindAllStringSubmatch(string(content), -1)
		if len(m) > 0 {
			// sm is an array with all (sub)matches. First element is the whole regex match, and the following
			// elements are the matches of the capturing groups.
			for _, sm := range m {
				match := sm[0]
				varName := sm[1]
				defaultValue := sm[2]
				var realValue string
				envValue := os.Getenv(varName)
				if envValue != "" {
					realValue = envValue
				} else {
					realValue = defaultValue
				}
				if c.GlobalBool("debug") || config.Debug {
					log.Printf("[%s] Found %s (default: %s) using: %s", file, varName, defaultValue, realValue)
				}
				content = bytes.Replace(content, []byte(match), []byte(realValue), -1)
			}
		}

		if !c.Bool("noop") {
			err = ioutil.WriteFile(file, content, 0644)
			if err != nil {
				return fmt.Errorf("error writing file %s: %v", file, err)
			}
		} else {
			if c.GlobalBool("debug") || config.Debug {
				log.Println("not saving file because of noop switch")
			}
		}

	}
	return nil
}
