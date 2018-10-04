package main

import (
	"flag"
)

type programParams struct {
	configPath string
	search string
	folder string
	count int
	excludes arrayFlags
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func getProgramParams() programParams {
	configPath := flag.String("cfg", "config.toml", "config file path")
	search := flag.String("search", "", "tag")
	folder := flag.String("folder", "", "name of the folder")
	searchCount := flag.Int("count", 0, "number of images")

	var excludeSearch arrayFlags
	flag.Var(&excludeSearch, "exclude", "Tags that should be excluded from search")
	flag.Parse()
	return programParams{
		*configPath,
		*search,
		*folder,
		*searchCount,
		excludeSearch,
	}
}