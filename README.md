# GoImageDownloader
GoImageDownloader is a simple, efficient command-line tool written in Go for downloading images from any given URL. It takes two parameters: the URL from which images will be downloaded and the local path where the images should be saved.

## Features
* Downloads all images from a specified URL.
* Saves images to a specified local directory, creating the directory if it doesn't exist.
* Uses GoQuery for parsing HTML and extracting image URLs.
## Requirements
* Go (1.15 or later recommended)
* GoQuery package
## Installation
To use GoImageDownloader, first ensure you have Go installed on your system. You can then download and install GoImageDownloader by cloning this repository and installing the required GoQuery package:

```bash
go get github.com/PuerkitoBio/goquery
```
## Usage
GoImageDownloader is easy to use with two simple command-line flags:

*`-url:` The URL to download images from (e.g., "https://unsplash.com/de/s/fotos/sample")
*`-path:` The local file system path where the images will be saved (e.g., "images")

Example:

```bash
go run main.go -url https://unsplash.com/de/s/fotos/sample -path ./images
```

This command downloads all images from the specified Unsplash page and saves them to the ./images directory.

## Contributing
Contributions to GoImageDownloader are welcome! Please feel free to submit pull requests or open issues to suggest improvements or report bugs.

## License
GoImageDownloader is open-source software licensed under the MIT license.