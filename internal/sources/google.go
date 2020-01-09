package imagesources

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/henrymxu/imagedownloader/internal/utils"
)

const googleBaseUrl = "https://www.google.co.in/search?q=%s&source=lnms&tbm=isch"

type google struct {

}

func NewGoogleSource() *google {
	return &google{}
}

func (g *google) GetImageUrls(count int, search string, excludes ...string) []string {
	url := fmt.Sprintf(googleBaseUrl, search)
	doc := utils.MakeRequestAndCreateDocument(url)
	// Find the review items
	imageUrls := make([]string, count)
	index := 0
	doc.Find(".rg_meta").Each(func(i int, s *goquery.Selection) {
		if index >= count {
			return
		}
		result := map[string]interface{}{}
		text := s.Contents().Text()
		_ = json.Unmarshal([]byte(text), &result)
		url := result["ou"].(string)
		imageUrls[index] = url
		index++
	})
	return imageUrls
}