package imagesources

import (
	"encoding/json"
	"fmt"
	"github.com/henrymxu/imagedownloader/internal/utils"
	"strconv"
	"strings"
)

const baseUrl = "https://api.flickr.com/services/rest/?"
const imageBaseUrl = "https://farm%d.staticflickr.com/%s/%s_%s.jpg"
const flickrMaxImages = 500

type flickr struct {
	params map[string]string
}

func NewFlickrSource(apiKey string) *flickr {
	params := make(map[string]string)
	params["api_key"] = apiKey
	params["method"] = "flickr.photos.search"
	params["format"] = "json"
	params["nojsoncallback"] = "1"
	params["per_page"] = strconv.Itoa(flickrMaxImages)
	return &flickr{params}
}

// Initial flickr rest call returns list of images with various properties including server, farm, id, etc.
// Image urls are then constructed using these properties
func (flickr *flickr) GetImageUrls(count int, search string, excludes ...string) []string {
	tags := strings.Builder{}
	if excludes != nil {
		for _, exclude := range excludes {
			tags.WriteString(exclude)
			tags.WriteString(",-")
		}
	}
	flickr.params["text"] = search
	flickr.params["tags"] = tags.String()
	var allImageUrls []string
	var pages = count/flickrMaxImages + 1
	for i := 1; i <= pages; i++ {
		flickr.params["page"] = strconv.Itoa(i)
		response := utils.MakeRequestWithQuery(baseUrl, flickr.params)
		allImageUrls = append(allImageUrls, parseHttpResponse(response)...)
	}
	return allImageUrls[:count]
}

func parseHttpResponse(response []byte) []string {
	var formatted map[string]interface{}

	err := json.Unmarshal(response, &formatted)
	if err != nil {
		panic(err)
	}
	imageResults := formatted["photos"].(map[string]interface{})["photo"].([]interface{})
	imageUrls := make([]string, len(imageResults))
	for index, photo := range imageResults {
		imageUrls[index] = buildImageUrl(photo.(map[string]interface{}))
	}
	return imageUrls
}

func buildImageUrl(image map[string]interface{}) string {
	return fmt.Sprintf(imageBaseUrl, int(image["farm"].(float64)), image["server"], image["id"], image["secret"])
}
