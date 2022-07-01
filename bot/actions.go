package bot

import (
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/vova616/screenshot"
	"github.com/pkg/browser"
)

func GetScreenshot() tgbotapi.FileBytes {
	filename := "img.png"

	img, imgErr := screenshot.CaptureScreen()	
	if (imgErr != nil) {
		panic(imgErr)
	}

	file, fileErr := os.Create(filename)
	if fileErr != nil {
		panic(fileErr)
	}
	defer file.Close()

	encodeErr := jpeg.Encode(file, img, nil)
	if encodeErr != nil {
		panic(encodeErr)
	}

	photoBytes, err := ioutil.ReadFile(filename)	
	if err != nil {
		panic(err)
	}
		
	return tgbotapi.FileBytes {
		Name:  "picture",
		Bytes: photoBytes,
	}
}

func KillProcess(name string) string {
    processes, err := process.Processes()
    if err != nil {
        panic(err)
    }
    for _, p := range processes {
        n, _ := p.Name()
        if n == name {
            err := p.Kill()
			if (err != nil) {
				return "Process not found"
			}
        } else {
			return "Process not found"
		}
    }
    return "Terminated!"
}

func GetAllProcesses() string {
	processes, err := process.Processes()
	var list []string
	m := make(map[string]string)

	if (err != nil) {
		panic(err)
	}

	for _, p := range processes {
		if (p.Foreground()); err == nil {
			name, _ := p.Name()
			m[name] = name
		}
	}

	for _, v := range m {
		list = append(list, v)
	}	

	return strings.Join(list, "\n")
}

func OpenBrowser(url string) {
	browser.OpenURL(url)
}