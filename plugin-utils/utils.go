package plugin_utils

import (
    "strings"
    "regexp"
    "platform-plugin/output"
)

const (
    pluginErrorRep = `rpc error: code = \w*\s* desc = `
)

func IsPluginMessage(string string) bool {
    return strings.Contains(string, output.StartToken)
}

func GetPluginError(err error) string {
    if err == nil {
        return ""
    }

    pat := regexp.MustCompile(pluginErrorRep)

    return pat.ReplaceAllString(strings.Trim(strings.Trim(err.Error(), "\n"), "\r"), "")
}
