package storage

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	FlickrApiKey string

	ImagesFolder string
	ImagesNameFormat string

	GoRoutineCount int
}

func SaveToDisc(contents []byte, path string) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)

	file, err := os.Create(path)
	check(err)

	err = ioutil.WriteFile(path, contents, 0644)
	check(err)
	file.Close()
}

func ReadFromDisc(path string) []byte {
	data, err := ioutil.ReadFile(path)
	check(err)
	return data
}

func MoveFile(originalPath string, newPath string) {
	err := os.Rename(originalPath, newPath)
	check(err)
}

func GetHomeDir() string {
	usr, err := user.Current()
	check(err)
	return usr.HomeDir
}

func LoadConfig(configPath string) Config {
	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	check(err)
	return conf
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
