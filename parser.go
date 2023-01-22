package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	"image/png"
	_ "image/png" // force import the png decoder
)

func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func cropScreenshotToElement(rawImage []byte, xStart int, yStart int, xEnd int, yEnd int, filename string) {
	// load pre image as image.image
	img, _, err := image.Decode(bytes.NewBuffer(rawImage))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(xStart)

	// crop image
	cropped, err := cropImage(img, image.Rect(xStart, yStart, xEnd, yEnd))
	if err != nil {
		log.Fatal(err)
	}

	// save cropped one
	out2, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out2.Close()

	png.Encode(out2, cropped)
}

func takeScreenshot(filename string, url string) {
	const (
		seleniumPath     = "./selenium-server/selenium-server-standalone-3.9.1.jar"
		chromeDriverPath = "./chromedriver_linux64/chromedriver"
		port             = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Stop()

	// connect to the WebDriver instance running locally
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless",
			"--no-sandbox",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatal(err)
	}
	defer wd.Quit()

	if err := wd.Get(url); err != nil {
		log.Fatal(err)
	}

	screenshot, err := wd.Screenshot()
	if err != nil {
		log.Fatal(err)
	}

	element, err := wd.FindElement(selenium.ByCSSSelector, ".cardapio")
	if err != nil {
		log.Fatal(err)
	}

	elementSize, err := element.Size()
	if err != nil {
		log.Fatal(err)
	}

	location, err := element.Location()
	if err != nil {
		log.Fatal(err)
	}

	heightMagicNumber := 65 + 30 + 37 - 2

	cropScreenshotToElement(screenshot, location.X, location.Y-heightMagicNumber, location.X+elementSize.Width, location.Y+elementSize.Height, filename)
}
