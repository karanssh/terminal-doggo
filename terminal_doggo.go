package main

import (
	"encoding/json"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	doggoData := fetchImage()
	for !strings.HasSuffix(strings.ToLower(doggoData.URL), "jpg") {
		log.Print("did not get jpeg oops, try again")
		doggoData = fetchImage()

	}

	//fetch image
	gotImage, err := http.Get(doggoData.URL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("got : %v", gotImage.Body)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, gotImage.Body)
	if err != nil {
		log.Fatal(err)
	}
	convertImage(filePath)

}

func fetchImage() doggoEndPointGet {
	//fetch image data
	resp, err := http.Get(doggoEndPoint)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var doggoData doggoEndPointGet

	err = json.Unmarshal(body, &doggoData)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("got : ", doggoData)
	return doggoData

}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

var (
	doggoEndPoint = "https://random.dog/woof.json?include=jpg"
	filePath      = "current.jpg"
)

type doggoEndPointGet struct {
	FileSizeBytes int    `json:"fileSizeBytes"`
	URL           string `json:"url"`
}
