package settings

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PATH string = "userinfo.txt"
const COMMANDS string = `
/help -> get a help
/screenshot -> take a screenshot of machine and send
/kill_proc -> to kill the process by its name
/browse -> enter url in http format to open it in the browser
`

var USERID int64
var TOKEN string

func Initialize() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Perss any key to continue... ")
	input, _ := reader.ReadString('\n')
	if strings.TrimSpace(input) == "credentials" {
		removeCredentials()
		Initialize()
	}

	if _, err := os.Stat(PATH); errors.Is(err, os.ErrNotExist) {
		fmt.Print("Enter your telegram user id: ")
		userId, _ := reader.ReadString('\n')
		fmt.Print("Enter your telegram bot api token: ")
		token, _ := reader.ReadString('\n')

		file, fileErr := os.Create(PATH)
		if fileErr != nil {
			panic(fileErr)
		}
		defer file.Close()

		_, ioErr := file.WriteString(strings.TrimSpace(userId) + "\n" + strings.TrimSpace(token))
		if ioErr != nil {
			panic(ioErr)
		}
	} else {
		file, fileErr := os.Open(PATH)
		if fileErr != nil {
			panic(fileErr)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lines := []string{}
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if scanerErr := scanner.Err(); scanerErr != nil {
			panic(scanerErr)
		}

		userId, convErr := strconv.ParseInt(lines[0], 10, 64)
		if convErr != nil {
			panic(convErr)
		}
		USERID = userId
		TOKEN = lines[1]
	}
}

func removeCredentials() {
	err := os.Remove(PATH)
    if err != nil {
        panic(err)
    }
}