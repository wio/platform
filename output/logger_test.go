package output

import (
    "testing"
    "bytes"
    "os"
    "log"
    "errors"
)

func TestInfo(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "INFO Hello INFO There INFO My Man"

    Info("Hello ")
    Info("There ")
    Info("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Info Output was incorrect, got:: %s, want:: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "INFO Hello\nINFO There\nINFO My Man\n"

    Info("Hello\n")
    Info("There\n")
    Info("My Man\n")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Info Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "INFO Hello\nThere\nMy Man\n45"

    Info("Hello\n%s\n%s\n%d", "There", "My Man", 45)

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Info Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }
}

func TestInfoln(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "INFO Hello \nINFO There \nINFO My Man\n"

    Infoln("Hello ")
    Infoln("There ")
    Infoln("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Info Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "INFO Hello\n\nINFO There\n\nINFO My Man\n\n"

    Infoln("Hello\n")
    Infoln("There\n")
    Infoln("My Man\n")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Info Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "INFO Hello\nThere\nMy Man\n\n"

    Infoln("Hello\n%s\n%s\n", "There", "My Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Info Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }
}

func TestWarn(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "WARN Hello WARN There WARN My Man"

    Warn("Hello ")
    Warn("There ")
    Warn("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Warn Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "WARN Hello\nWARN There\nWARN My Man\n"

    Warn("Hello\n")
    Warn("There\n")
    Warn("My Man\n")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Warn Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "WARN Hello\nThere\nMy Man\n"

    Warn("Hello\n%s\n%s\n", "There", "My Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Warn Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }
}

func TestWarnln(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "WARN Hello \nWARN There \nWARN My Man\n"

    Warnln("Hello ")
    Warnln("There ")
    Warnln("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Warn Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "WARN Hello\n\nWARN There\n\nWARN My Man\n\n"

    Warnln("Hello\n")
    Warnln("There\n")
    Warnln("My Man\n")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Warn Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "WARN Hello\nThere\nMy Man\n\n"

    Warnln("Hello\n%s\n%s\n", "There", "My Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Warn Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }
}

func TestFatal(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "ERR Hello ERR There ERR My Man"

    Fatal("Hello ")
    Fatal("There ")
    Fatal("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "ERR Hello\nERR There\nERR My Man\n"

    Fatal("Hello\n")
    Fatal("There\n")
    Fatal("My Man\n")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "ERR Hello\nThere\nMy Man\n"

    Fatal("Hello\n%s\n%s\n", "There", "My Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Error Test ///////////
    desiredOutput = "ERR this is an error"

    Fatal(errors.New("this is an error"))

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }
}

func TestFatalln(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "ERR Hello \nERR There \nERR My Man\n"

    Fatalln("Hello ")
    Fatalln("There ")
    Fatalln("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "ERR Hello\n\nERR There\n\nERR My Man\n\n"

    Fatalln("Hello\n")
    Fatalln("There\n")
    Fatalln("My Man\n")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "ERR Hello\nThere\nMy Man\n\n"

    Fatalln("Hello\n%s\n%s\n", "There", "My Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Error Test ///////////
    desiredOutput = "ERR this is an error\n"

    Fatalln(errors.New("this is an error"))

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Fatal Output was incorrect, got: %s, want: %s", decodedMessage, desiredOutput)
    }
}
