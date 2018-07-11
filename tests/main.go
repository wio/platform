package main

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/hashicorp/go-plugin"
    "plugins/shared"
    "bytes"
    "io"
    "strings"
    "plugins/plugin-utils"
    "io/ioutil"
    "path/filepath"
    "github.com/hashicorp/go-hclog"
    "plugins/output"
    "encoding/json"
)

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{}

type Platform struct {
    Name string
    Exec string
}

func main() {
    ///////// READ Plugin name from platform.json file /////////

    wioPlatform := &Platform{}

    execPath, err := os.Executable()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }

    execPath, err = filepath.Abs(execPath + "/../")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }

    buff, err := ioutil.ReadFile(execPath + "/avr-platform/platform.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }

    err = json.Unmarshal(buff, wioPlatform)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }

    //////////////////////////////////////////////////////////////

    /////////////////////// Create Client ///////////////////////

    old := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // create a logger
    logger := hclog.New(&hclog.LoggerOptions{
        Output: os.Stdout,
        Level:  hclog.Trace,
        Name:   "wio",
    })

    PluginMap[wioPlatform.Name] = &shared.PlatformPlugin{}

    // We're a host. Start by launching the plugin process.
    client := plugin.NewClient(&plugin.ClientConfig{
        HandshakeConfig: shared.Handshake,
        Plugins:         PluginMap,
        Logger: logger,
        Cmd:             exec.Command(execPath + "/avr-platform/" + wioPlatform.Exec),
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

    ////////////////////////////////////////////////////////////

    ///////////////////////// Get Plugin ///////////////////////

    // Request the plugin
    raw, err := rpcClient.Dispense("platform-atmelavr")
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }

    platform := raw.(shared.Platform)

    ///////////////////////////////////////////////////////////

    //////////////////////// Call functions ///////////////////

    obj := &shared.TargetInformation{}
    obj.Platform = "Hola"

    // build function
    executables, err := platform.BuildProject(obj)

    fmt.Printf("wio message:: {{Number of executables given %d\n}}", len(executables))
    if err != nil {
       fmt.Fprintln(old, plugin_utils.GetPluginError(err))
       os.Exit(2)
    }

    // run function
    _, err = platform.RunProject(&shared.RunInformation{})
    if err != nil {
        fmt.Fprintln(old, plugin_utils.GetPluginError(err))
        os.Exit(2)
    }

    // print the output because we need to structure the logs
    w.Close()
    os.Stdout = old

    var buf bytes.Buffer
    io.Copy(&buf, r)

    for _, line := range strings.Split(buf.String(), "\n") {
        if plugin_utils.IsPluginMessage(line) {
            fmt.Println(output.DecodeMessage(line))
        }
    }
}
