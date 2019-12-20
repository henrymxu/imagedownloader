package main

import (
	"fmt"
	"github.com/henrymxu/imagedownloader"
	"github.com/henrymxu/imagedownloader/imgsrcs"
	"github.com/henrymxu/imagedownloader/utils/program"
	"github.com/henrymxu/imagedownloader/utils/storage"
	"strings"
)

// params: config file path, search, folder name, number of images, tags to exclude ...
func main() {
	params := program.GetInitialParams()

	configPath := params.ConfigPath
	search := params.Search
	folder := params.Folder
	count := params.Count
	excludesearch := params.Excludes

	config := storage.LoadConfig(configPath)
	goroutineCount := config.GoRoutineCount

	// Create an ImageSource
	var imageSource imagesources.ImageSource
	imageSource = imagesources.New(config.FlickrApiKey)

	// Configure path from parameters
	imageNameFormat := config.ImagesNameFormat
	if !strings.Contains(imageNameFormat, "%d") {
		imageNameFormat = fmt.Sprintf("%s_%%d.jpg", imageNameFormat)
	}
	imageFolder := fmt.Sprintf("%s/%s", config.ImagesFolder, folder)
	path := fmt.Sprintf("%s/%s/%s", storage.GetHomeDir(), imageFolder, imageNameFormat)

	// Run Image Downloader
	var imgdlr = imagedownloader.New(imageSource, goroutineCount)
	imgdlr.DownloadImages(path, count, search, excludesearch)
}

