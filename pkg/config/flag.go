package config

import (
	"flag"
	"os"
)

func UpdateConfigWithFlag(c *Config) {
	var flags flag.FlagSet
	flags.BoolVar(&c.Verbose, "verbose", false, "Verbose Display")
	flags.Parse(os.Args[1:])
}
