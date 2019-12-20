# Image Downloader

Program to download large amounts of images from public image sources

## Currently supported Image sources

- Flickr

## Example command line usage

`./imagedownloader -cfg=config.toml -source=flickr "-search=porsche 918" -folder=~/Documents/ImageSet/porsche918 -format=image_%d.jpg -count=250 -exclude=911 -exclude=turbo`

- cfg = location of config file, follow the template provided in `_examples/config_example.toml`
- source = source of images, requires correct api key to be found in `config.tml` if source requires API key (default is `flickr`)
- search = keyword for images
- folder = name of the folder where images will be saved, path is absolute (can use `~`)
- format = format for names of images, if `%s` is not provided, `_%s` will be appended to the end, where %s is the image number (default is `image_%d.jpg`)
- count = number of images to download (default is `250`)
- exclude (optional) = keywords that should be excluded from search

## Config
- FlickrApiKey = ApiKey required to use Flickr services

# License

MIT