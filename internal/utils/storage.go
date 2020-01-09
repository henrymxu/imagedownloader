package utils

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	FlickrApiKey string
}

func BuildPathFromParams(imageFormat string, imageFolder string)string  {
	// Configure path from parameters
	imageNameFormat := imageFormat
	if !strings.Contains(imageNameFormat, "%d") {
		imageNameFormat = fmt.Sprintf("%s_%%d.jpg", imageNameFormat)
	}
	if strings.Contains(imageFolder, "~") {
		imageFolder = strings.Replace(imageFolder, "~", GetHomeDir(), -1)
	}
	path := fmt.Sprintf("%s/%s", imageFolder, imageNameFormat)
	return path
}

func SaveToDisc(contents []byte, path string) bool {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)

	file, err := os.Create(path)
	check(err)

	err = ioutil.WriteFile(path, contents, 0644)
	if err != nil {
		return false
	}
	err = file.Close()
	if err != nil {
		return false
	}
	return true
}

func GetHomeDir() string {
	usr, err := user.Current()
	check(err)
	return usr.HomeDir
}

func LoadConfig(configPath string) *Config {
	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	check(err)
	return &conf
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
