package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vova616/screenshot"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func GetScreenshot() tgbotapi.FileBytes {
	img, err := screenshot.CaptureScreen()
	
	if (err == nil) {
		f, err := os.Create("img.jpg")

		if err != nil {
			panic(err)
		}

		defer f.Close()

		err = jpeg.Encode(f, img, nil)

		if err != nil {
			panic(err)
		}
	}

	photoBytes, err := ioutil.ReadFile("img.jpg")
		
	if err != nil {
		panic(err)
	}
		
	photoFileBytes := tgbotapi.FileBytes {
		Name:  "picture",
		Bytes: photoBytes,
	}

	return photoFileBytes
}