package output

import (
    "testing"
    "bytes"
    "log"
    "os"
)

func TestEncode(t *testing.T) {
    desired1 := "wio message:: {{Hello There}}"
    desired2 := "wio message:: {{Hello\nThere}}"

    encoded1 := encodeMessage("Hello There")
    encoded2 := encodeMessage("Hello\nThere")

    if encoded1 != desired1 {
        t.Errorf("Encoding was incorrect, got %s, want %s", encoded1, desired2)
    } else if encoded2!= desired2 {
        t.Errorf("Encoding was incorrect, got %s, want %s", encoded1, desired2)
    }
}

func TestDecodeMessage(t *testing.T) {
    desired1 := "Hello Brother"
    desired2 := "Hello \nBrother. No way"

    encoded1 := encodeMessage(desired1)
    encoded2 := encodeMessage(desired2)

    decoded1 := DecodeMessage(encoded1)
    decoded2 := DecodeMessage(encoded2)

    if decoded1 != desired1 {
        t.Errorf("Decoding was incorrect, got %s, want %s", decoded1, desired2)
    } else if decoded2 != decoded2 {
        t.Errorf("Decoding was incorrect, got %s, want %s", decoded2, desired2)
    }
}

func TestPrint(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()

    /////// Print Test ///////////
    desiredOutput := "Hello There My Man"

    Print("Hello ")
    Print("There ")
    Print("My Man")

    decodedMessage := DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Output was incorrect, got %s, want %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Println Test ///////////
    desiredOutput = "Hello\nThere My\nMan"

    Println("Hello")
    Println("There My")
    Print("Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Output was incorrect, got %s, want %s", decodedMessage, desiredOutput)
    }

    buf.Reset()

    /////// Printf Test ///////////
    desiredOutput = "Hello\nThere My\nMan"

    Printf("Hello\n%s %s\n%s", "There", "My", "Man")

    decodedMessage = DecodeMessage(buf.String())

    if decodedMessage != desiredOutput {
        t.Errorf("Output was incorrect, got %s, want %s", decodedMessage, desiredOutput)
    }
}

