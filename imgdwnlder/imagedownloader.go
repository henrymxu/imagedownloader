package imgdwnlder

import (
	"fmt"
	"github.com/henrymxu/imagedownloader/imgsrcs"
	"github.com/henrymxu/imagedownloader/utils"
	"sync"
	"time"
)

type ImageDownloader struct {
	source         imagesources.ImageSource
	goroutineCount int
}

func New(source imagesources.ImageSource, goroutineCount int) ImageDownloader {
	return ImageDownloader{source, goroutineCount}
}

func (imgdlr ImageDownloader) DownloadImages(path string, count int, search string, excludes []string) {
	urls := imgdlr.source.GetImageUrls(count, search, excludes...)
	fmt.Printf("%d image urls discovered\n", len(urls))
	imgdlr.downloadAllImages(path, urls)
}

func (imgdlr ImageDownloader) downloadAllImages(path string, imageurls []string) {
	start := time.Now()
	var imageUrlMaps = imgdlr.splitUpImageUrls(imgdlr.goroutineCount, imageurls)
	var wg sync.WaitGroup
	images := 0
	for i := 0; i < imgdlr.goroutineCount; i++ {
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
func (imgdlr ImageDownloader) splitUpImageUrls(count int, imageUrls []string) []map[string]int {
	if imgdlr.goroutineCount > count {
		imgdlr.goroutineCount = count
	}
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
