package main

import (
	"fmt"
	"github.com/henrymxu/imagedownloader/imgsrcs"
	"github.com/henrymxu/imagedownloader/utils/network"
	"github.com/henrymxu/imagedownloader/utils/storage"
	"strings"
	"sync"
	"time"
)

var count int
var goroutineCount int

var search string
var folder string
var excludeSearch arrayFlags

// params: config file path, search, folder name, number of images, tags to exclude ...
func main() {
	params := getProgramParams()

	configPath := params.configPath
	search = params.search
	folder = params.folder
	count = params.count
	excludeSearch = params.excludes

	config := storage.LoadConfig(configPath)
	goroutineCount = config.GoRoutineCount

	var imageSource imagesources.ImageSource
	imageSource = imagesources.New(config.FlickrApiKey)
	imageUrls := imageSource.GetImageUrls(count, search, excludeSearch...)

	imageNameFormat := config.ImagesNameFormat
	if !strings.Contains(imageNameFormat, "%d") {
		imageNameFormat = fmt.Sprintf("%s_%%d.jpg", imageNameFormat)
	}
	imageFolder := fmt.Sprintf("%s/%s", config.ImagesFolder, folder)

	path := fmt.Sprintf("%s/%s/%s", storage.GetHomeDir(), imageFolder, imageNameFormat)
	downloadAllImages(path, imageUrls)
}

func downloadAllImages(path string, imageUrls []string) {
	start := time.Now()
	var imageUrlMaps = splitUpImageUrls(goroutineCount, imageUrls)
	var wg sync.WaitGroup
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		index := i
		go func() {
			defer wg.Done()
			for key, value := range imageUrlMaps[index] {
				imageData := network.MakeSimpleRequest(key)
				storage.SaveToDisc(imageData, fmt.Sprintf(path, value))
			}
		}()
	}
	wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// This function splits up the list of image urls into smaller maps so a single goroutine can service multiple images
// Returns a map with key image url, and value of original index
func splitUpImageUrls(count int, imageUrls []string) []map[string]int {
	if goroutineCount > count {
		goroutineCount = count
	}
	var imageUrlMaps = make([]map[string]int, count)
	for i := 0; i < len(imageUrls); i += count {
		for j := i; j < i + count; j++ {
			if imageUrlMaps[j - i] == nil {
				imageUrlMaps[j - i] = make(map[string]int)
			}
			imageUrlMaps[j - i][imageUrls[j]] = j
		}
	}
	return imageUrlMaps
}
