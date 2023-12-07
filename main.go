package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func saveMusic(songName, url string) error {
	if strings.Contains(songName, "/") {
		s := strings.Split(songName, "/")
		songName = s[0] + s[1]
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	// Create the file
	out, err := os.Create("/home/arxonic/Desktop/face/" + songName + ".mp3")
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func parseMovies(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("li.item").Each(func(i int, s *goquery.Selection) {
		var songName = s.Find("span.track").Text()
		var url, _ = s.Find("li.play").Attr("data-url")
		i++
		fmt.Println(i, "\t", songName, "\t", url)
		// err := saveMusic(songName, url)
		// if err != nil {
		// 	fmt.Println("Error downloading file: ", err)
		// 	return
		// }

		// fmt.Println("Downloaded: " + url)
	})
}

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://muzofond.fm/collections/artists/face/", "https://muzofond.fm/collections/artists/face/2"},
		ParseFunc: parseMovies,
	}).Start()
}
