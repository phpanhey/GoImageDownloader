package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	urlArg := flag.String("url", "https://unsplash.com/de/s/fotos/sample", "a url to download images from")
	pathArg := flag.String("path", "images", "a path to save images")
	flag.Parse()

	html, err := downloadHtml(*urlArg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	imagesUrls := extractImageUrls(html)
	donwloadImages(imagesUrls, *pathArg)

}
func downloadHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "Error Downloading html", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response body", err
	}

	return string(body), nil
}

func extractImageUrls(html string) []string {
	imagesUrls := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return imagesUrls
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imagesUrls = append(imagesUrls, s.AttrOr("src", ""))
	})
	return imagesUrls
}

func donwloadImages(imagesUrls []string, path string) {
	createPathIfNotExists(path)
	for _, url := range imagesUrls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error downloading image: ", err)
			continue
		}
		defer resp.Body.Close()

		image, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading image: ", err)
			continue
		}

		err = os.WriteFile(path+"/"+url[strings.LastIndex(url, "/")+1:], image, 0666)
		if err != nil {
			fmt.Println("Error writing image: ", err)
			continue
		}
	}
}

func createPathIfNotExists(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
