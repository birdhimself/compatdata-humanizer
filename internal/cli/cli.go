package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var titleColor = color.New(color.Bold, color.Underline).PrintfFunc()
var successColor = color.New(color.FgGreen).PrintfFunc()
var errorColor = color.New(color.FgRed, color.Bold).PrintfFunc()
var warningColor = color.New(color.FgYellow).PrintfFunc()
var confirmColor = color.New(color.FgCyan, color.Bold).PrintfFunc()
var infoColor = color.New(color.FgBlue).PrintfFunc()

func Confirm(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	confirmColor("\n%s? [y/N]: ", question)

	res, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	res = strings.ToLower(strings.TrimSpace(res))

	return res == "y" || res == "yes"
}

func Title(title string) {
	titleColor("\n%s\n", title)
}

func Warning(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)

	warningColor("  [!] %s\n", message)
}

func Error(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)

	errorColor("  [!] %s\n", message)
}

func Success(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)

	successColor("  [+] %s\n", message)
}

func Info(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)

	infoColor("  [i] %s\n", message)
}

func BulletList(input []string) {
	out := ""

	for _, v := range input {
		out += fmt.Sprintf("  - %s\n", v)
	}

	fmt.Print(out)
}
