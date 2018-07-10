package main

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/hashicorp/go-plugin"
    "plugins/shared"
    "github.com/hashicorp/go-hclog"
)

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
    "avr-platform": &shared.PlatformPlugin{},
}


func main() {
    logger := hclog.New(&hclog.LoggerOptions{
        Output: os.Stdout,
        Level:  hclog.Trace,
        Name:   "wio",
    })

    // We're a host. Start by launching the plugin process.
    client := plugin.NewClient(&plugin.ClientConfig{
        HandshakeConfig: shared.Handshake,
        Plugins:         PluginMap,
        Logger:          logger,
        Cmd:             exec.Command("./avr-platform/avr"),
        AllowedProtocols: []plugin.Protocol{
            plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
        SyncStdout: os.Stdout,
    })
    defer client.Kill()

    // Connect via RPC
    rpcClient, err := client.Client()
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }

    // Request the plugin
    raw, err := rpcClient.Dispense("avr-platform")
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }

    platform := raw.(shared.Platform)
    obj := &shared.TargetInformation{}
    obj.Platform = "Hola"
    executables, err := platform.BuildProject(obj)
    fmt.Println(len(executables))
    if err != nil {
        fmt.Println(err.Error())
    }
}
