package output

import (
    "log"
    "fmt"
    "strings"
)

const (
    StartToken = "wio message:: {{"
    EndToken = "}}"
)

func encodeMessage(string string) string {
    return StartToken + string + EndToken
}

func DecodeMessage(givenMessage string) string {
    strToReturn := ""

    for _, line := range strings.Split(givenMessage, "}}") {
        if len(strings.Trim(line, "\n")) == 0 {
            continue
        }

        strToReturn += line[strings.Index(line, StartToken) + len(StartToken):]
    }

    return strToReturn
}

func Print(message string) {
    log.Println(encodeMessage(message))
}

func Println(message string) {
    Print(message + "\n")
}

func Printf(message string, a ...interface{}) {
    Print(fmt.Sprintf(message, a...))
}
