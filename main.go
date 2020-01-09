package main

import (
	"fmt"
	imagedownloader "github.com/henrymxu/imagedownloader/downloader"
	imagesources "github.com/henrymxu/imagedownloader/internal/sources"
	"github.com/henrymxu/imagedownloader/internal/utils"
)

// params: config file path, search, folder name, number of images, tags to exclude ...
func main() {
	params := utils.GetInitialParams()
	if params == nil {
		return
	}

	config := utils.LoadConfig(params.ConfigPath)

	// Create an ImageSource
	var imageSource imagesources.ImageSource
	switch params.Source {
	case "flickr":
		imageSource = imagesources.NewFlickrSource(config.FlickrApiKey)
	case "google":
		imageSource = imagesources.NewGoogleSource()
	}

	if imageSource == nil {
		fmt.Printf("Invalid Image Source (%s)\n", params.Source)
		return
	}

	path := utils.BuildPathFromParams(params.ImageFormat, params.Folder)
	imagedownloader.DownloadImages(imageSource, path, params.Count, params.Search, params.Excludes)
}
