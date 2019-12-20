package main

import (
	"github.com/henrymxu/imagedownloader/imgdwnlder"
	"github.com/henrymxu/imagedownloader/imgsrcs"
	"github.com/henrymxu/imagedownloader/utils"
)

// params: config file path, search, folder name, number of images, tags to exclude ...
func main() {
	params := utils.GetInitialParams()
	if params == nil {
		return
	}

	config := utils.LoadConfig(params.ConfigPath)
	if config == nil {
		return
	}

	// Create an ImageSource
	var imageSource imagesources.ImageSource
	imageSource = imagesources.New(config.FlickrApiKey)

	path := utils.BuildPathFromParams(params.ImageFormat, params.Folder)

	// Run Image Downloader
	var imgdlr = imgdwnlder.New(imageSource, 50)
	imgdlr.DownloadImages(path, params.Count, params.Search, params.Excludes)
}

