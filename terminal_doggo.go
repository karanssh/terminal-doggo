package main

import (
	"image"
	"log"
	"os"

	"github.com/qeesung/image2ascii/convert"
)

func main() {

	//fetch image data
	// resp, err := http.Get(doggoEndPoint)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// var doggoData doggoEndPointGet

	// err = json.Unmarshal(body, &doggoData)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Print("got : ", doggoData)
	// if !strings.HasSuffix(doggoData.URL, "jpg") {
	// 	log.Print("did not get jpeg oops")
	// 	os.Exit(1)
	// }

	// //fetch image
	// gotImage, err := http.Get(doggoData.URL)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("got : %v", gotImage.Body)
	// file, err := os.Create("asdf.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// _, err = io.Copy(file, gotImage.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	img, err := getImageFromFilePath("asdf.jpg")
	if err != nil {
		log.Fatal(err)
	}
	converter := convert.NewImageConverter()
	// dogImage, _, err := image.Decode(img)
	stringData := converter.Image2CharPixelMatrix(img, &convert.DefaultOptions)
	log.Print(stringData)
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
	doggoEndPoint = "https://random.dog/woof.json"
)

type doggoEndPointGet struct {
	FileSizeBytes int    `json:"fileSizeBytes"`
	URL           string `json:"url"`
}
