# Image Downloader

Program to download large amounts of images from public image sources

## Currently supported Image sources

- Flickr

## Example command line usage

`./imagedownloader -cfg=config.toml "-search=porsche 918" -folder=~/Documents/ImageSet/porsche918 -format=image_%d.jpg -count=250 -exclude=911 -exclude=turbo`

- cfg = location of config file, follow the template provided in `_examples/config_example.toml`
- search = keyword for images
- folder = name of the folder where images will be saved, path is absolute (can use `~`)
- format = format for names of images, if `%s` is not provided, `_%s` will be appended to the end, where %s is the image number
- count = number of images to download
- exclude (optional) = keywords that should be excluded from search

## Config
- FlickrApiKey = ApiKey required to use Flickr services

# License

MIT