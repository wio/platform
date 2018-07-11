package output

import (
    "fmt"
)

// Log type levels
type Type int

const (
    INFO      Type = 0
    WARN      Type = 1
    ERR       Type = 2
)

var logTypeTags = [3]string{"INFO", "WARN", "ERR"}

// Generic Write function
func write(args ...interface{}) {
    logType := INFO
    message := ""
    newline := false
    printfArgs := make([]interface{}, 0, len(args))

    for _, arg := range args {
        switch val := arg.(type) {
        case Type:
            logType = val
            break
        case string:
            if "" == message {
                message = val
            } else {
                printfArgs = append(printfArgs, val)
            }
            break
        case bool:
            newline = val
            break
        case error:
            message = val.Error()
            break
        default:
            printfArgs = append(printfArgs, val)
            break
        }
    }
    if newline {
        message = message + "\n"
    }

    Printf("%s %s", logTypeTags[logType], fmt.Sprintf(message, printfArgs...))
    if newline {
        println("")
    }
}

func Info(args ...interface{}) {
    write(append(args, INFO)...)
}

func Infoln(args ...interface{}) {
    write(append(args, INFO, true)...)
}

func Warn(args ...interface{}) {
    write(append(args, WARN)...)
}

func Warnln(args ...interface{}) {
    write(append(args, WARN,  true)...)
}

func Fatal(args ...interface{}) {
    write(append(args, ERR)...)
}

func Fatalln(args ...interface{}) {
    write(append(args, ERR, true)...)
}
