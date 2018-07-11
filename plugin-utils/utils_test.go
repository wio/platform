package plugin_utils

import (
    "testing"
    "github.com/go-errors/errors"
)

func TestIsPluginMessage(t *testing.T) {
    validMessage := "rpc codkjjk wio message:: {{Unknown desc = Some Error occurred}}"

    invalidMessag1 := "rpc codkjjk wio message: {{Unknown desc = Some Error occurred}}"
    invalidMessag2 := "rpc codkjjk wiomessage:: {{Unknown desc = Some Error occurred"
    invalidMessag3 := "rpc codkjjk message:: {{Unknown desc = Some Error occurred"

    check1 := IsPluginMessage(validMessage)
    check2 := IsPluginMessage(invalidMessag1)
    check3 := IsPluginMessage(invalidMessag2)
    check4 := IsPluginMessage(invalidMessag3)

    if !check1 {
        t.Errorf("(Valid check) Plugin Message check failed, for %s", validMessage)
    } else if check2 {
        t.Errorf("(Invalid check) Plugin Message check failed, for %s", invalidMessag1)
    } else if check3 {
        t.Errorf("(Invalid check) Plugin Message check failed, for %s", invalidMessag2)
    } else if check4 {
        t.Errorf("(Invalid check) Plugin Message check failed, for %s", invalidMessag3)
    }
}

func TestGetPluginError(t *testing.T) {
    error1 := errors.New("rpc error: code = Unknown desc = Some Error occurred")
    error2 := errors.New("rpc error: code = 2 desc = Some Error occurred")
    error3 := errors.New("rpc error: code = 45 desc = ")

    desired1 := "Some Error occurred"
    desired2 := "Some Error occurred"
    desired3 := ""

    decoded1 := GetPluginError(error1)
    decoded2 := GetPluginError(error2)
    decoded3 := GetPluginError(error3)

    if decoded1 != desired1 {
        t.Errorf("Error Output invalid, got: %s, want: %s", decoded1, desired1)
    } else if decoded2 != desired2 {
        t.Errorf("Error Output invalid, got: %s, want: %s", decoded2, desired2)
    } else if decoded3 != desired3 {
        t.Errorf("Error Output invalid, got: %s, want: %s", decoded3, desired3)
    }
}
