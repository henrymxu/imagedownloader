package imagedownloader

import (
	"fmt"
	"github.com/henrymxu/imagedownloader/imgsrcs"
	"github.com/henrymxu/imagedownloader/utils/network"
	"github.com/henrymxu/imagedownloader/utils/storage"
	"sync"
	"time"
)

type imagedownloader struct {
	source imagesources.ImageSource
	goroutineCount int
}

func New(source imagesources.ImageSource, goroutineCount int) imagedownloader {
	return imagedownloader{source, goroutineCount}
}

func (imgdlr imagedownloader) DownloadImages(path string, count int, search string, excludes []string) {
	urls := imgdlr.source.GetImageUrls(count, search, excludes...)
	imgdlr.downloadAllImages(path, urls)
}

func (imgdlr imagedownloader) downloadAllImages(path string, imageurls []string) {
	start := time.Now()
	var imageUrlMaps = imgdlr.splitUpImageUrls(imgdlr.goroutineCount, imageurls)
	var wg sync.WaitGroup
	for i := 0; i < imgdlr.goroutineCount; i++ {
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
func (imgdlr imagedownloader) splitUpImageUrls(count int, imageUrls []string) []map[string]int {
	if imgdlr.goroutineCount > count {
		imgdlr.goroutineCount = count
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