package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vova616/screenshot"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func GetScreenshot() tgbotapi.FileBytes {
	img, imgErr := screenshot.CaptureScreen()	
	if (imgErr != nil) {
		panic(imgErr)
	}

	file, fileErr := os.Create("img.jpg")
	if fileErr != nil {
		panic(fileErr)
	}
	defer file.Close()

	encodeErr := jpeg.Encode(file, img, nil)
	if encodeErr != nil {
		panic(encodeErr)
	}

	photoBytes, err := ioutil.ReadFile("img.jpg")	
	if err != nil {
		panic(err)
	}
		
	return tgbotapi.FileBytes {
		Name:  "picture",
		Bytes: photoBytes,
	}
}