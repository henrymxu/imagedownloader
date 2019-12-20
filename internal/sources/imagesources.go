package imagesources

// Interface for image sources such as google, flickr, etc.
// Objects that implement this should follow the Constructor Pattern for any initialization requirements.
// An ImageUrl refers to the exact url of the image (should be able to save response directly as image)

type ImageSource interface {
	GetImageUrls(count int, search string, excludes ...string) []string
}
