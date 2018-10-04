#Image Downloader

Program to download large amounts of images from public image sources

#Currently supported Image sources

- Flickr

#Example usage

./imagedownloader -cfg=config/config.toml "-search=porsche 918" -folder=porsche918 -count=250 -exclude=911 -exclude=turbo

- cfg = location of config file, follow the template provided in `config_example.toml`
- search = keyword for images
- folder = name of the folder where images will be saved (path is set in config file)
- count = number of images to download
- exclude = keywords that should be excluded from search

#Config
- FlickrApiKey = ApiKey required to use Flickr services
- ImagesFolder = Desired path relative to home directory to hold folders (Do not include home directory in path)
- ImagesNameFormat = Format for names of images, if `%s` is not provided, `_%s` will be appended to the end
- GoRoutineCount = Number of goroutines created for http requests

#License

MIT