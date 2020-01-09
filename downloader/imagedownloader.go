package imagedownloader

import (
	"fmt"
	"github.com/henrymxu/imagedownloader/internal/sources"
	"github.com/henrymxu/imagedownloader/internal/utils"
	"sync"
	"time"
)

const goroutineCount = 50

func DownloadImages(imgSource imagesources.ImageSource, path string, count int, search string, excludes []string) {
	urls := imgSource.GetImageUrls(count, search, excludes...)
	fmt.Printf("%d image urls discovered\n", len(urls))
	downloadAllImages(path, urls)
}

func downloadAllImages(path string, imageurls []string) {
	start := time.Now()
	var imageUrlMaps = splitUpImageUrls(goroutineCount, imageurls)
	var wg sync.WaitGroup
	images := 0
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		index := i
		go func() {
			defer wg.Done()
			for key, value := range imageUrlMaps[index] {
				imageData := utils.MakeSimpleRequest(key)
				if imageData != nil {
					success := utils.SaveToDisc(imageData, fmt.Sprintf(path, value))
					if success {
						images++
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("%.2fs elapsed, %d images downloaded\n", time.Since(start).Seconds(), images)
}

// This function splits up the list of image urls into smaller maps so a single goroutine can service multiple images
// Returns a map with key image url, and value of original index
func splitUpImageUrls(count int, imageUrls []string) []map[string]int {
	var imageUrlMaps = make([]map[string]int, count)
	for i := 0; i < len(imageUrls); i += count {
		for j := i; j < i+count; j++ {
			if imageUrlMaps[j-i] == nil {
				imageUrlMaps[j-i] = make(map[string]int)
			}
			imageUrlMaps[j-i][imageUrls[j]] = j
		}
	}
	return imageUrlMaps
}
