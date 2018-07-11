package plugin_utils

import (
    "strings"
    "plugins/output"
)

const (
    pluginError = "rpc error: code = Unknown desc = "
)

func IsPluginMessage(string string) bool {
    return strings.Contains(string, output.StartToken)
}

func GetPluginError(err error) string {
    return strings.Replace(strings.Trim(strings.Trim(strings.Trim(err.Error(), " "), "\n"), "\r"),
        pluginError, "", -1)
}
