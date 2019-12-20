package utils

import (
	"flag"
	"fmt"
)

type Params struct {
	ConfigPath  string
	Source      string
	Search      string
	Folder      string
	ImageFormat string
	Count       int
	Excludes    FlagArray
}

type FlagArray []string

func (i *FlagArray) String() string {
	return "Flag Array"
}

func (i *FlagArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func GetInitialParams() *Params {
	configPath := flag.String("cfg", "config.toml", "config file path")
	source := flag.String("source", "flickr", "source for images (flickr)")
	search := flag.String("search", "", "search keyword")
	folder := flag.String("folder", "", "name of the folder")
	imageFormat := flag.String("format", "image_%d.jpg", "image file format name")
	searchCount := flag.Int("count", 250, "number of images")

	var excludeSearch FlagArray
	flag.Var(&excludeSearch, "exclude", "Tags that should be excluded from search")
	flag.Parse()

	if *search == "" {
		fmt.Println("No search provided, please provide a search")
		return nil
	}

	if *folder == "" {
		fmt.Println("No folder provided, please provide a folder")
		return nil
	}

	return &Params{
		*configPath,
		*source,
		*search,
		*folder,
		*imageFormat,
		*searchCount,
		excludeSearch,
	}
}
