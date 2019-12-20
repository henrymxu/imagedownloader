package utils

import (
	"flag"
)

type Params struct {
	ConfigPath string
	Search     string
	Folder     string
	Count      int
	Excludes   FlagArray
}

type FlagArray []string

func (i *FlagArray) String() string {
	return "Flag Array"
}

func (i *FlagArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func GetInitialParams() Params {
	configPath := flag.String("cfg", "config.toml", "config file path")
	search := flag.String("search", "", "tag")
	folder := flag.String("folder", "", "name of the folder")
	searchCount := flag.Int("count", 0, "number of images")

	var excludeSearch FlagArray
	flag.Var(&excludeSearch, "exclude", "Tags that should be excluded from search")
	flag.Parse()
	return Params{
		*configPath,
		*search,
		*folder,
		*searchCount,
		excludeSearch,
	}
}
