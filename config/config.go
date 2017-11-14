package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"

	yaml "gopkg.in/yaml.v2"
	"github.com/urfave/cli"
)

// Values represents all configuration options that can be set in enver.conf
type Values struct {
	Files []string `yaml:"files"`
	Quiet bool     `yaml:"quiet"`
	Debug bool     `yaml:"debug"`
}

// FromFile loads the configurations settings into a Values struct, from a
// textfile.
func FromFile(c *cli.Context) (Values, error) {
	var config Values
	file := c.GlobalString("config")
	if file != "" {
		if _, notFoundFileErr := os.Stat(file); notFoundFileErr != nil {
			if c.GlobalBool("debug") {
				log.Printf("config file %s not found", file)
				return config, nil
			}
		} else {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return config, fmt.Errorf("error while reading configuration file: %v", err)
			}
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				return config, fmt.Errorf("error while parsing configuration file: %v", err)
			}
		}
	}
	return config, nil
}

// FilesToParse returns the list of files that we need to parse for variable replacement.
func (config Values) FilesToParse(c *cli.Context) ([]string, error) {
	var files []string

	// files passed as arguments take precedence over files specified in config file
	if len(c.StringSlice("files")) > 0 {
		files = c.StringSlice("files")
	} else if len(config.Files) > 0 {
		files = config.Files
	} else {
		return nil, fmt.Errorf("no files specified")
	}
	return files, nil
}
