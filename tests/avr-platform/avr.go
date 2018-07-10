package main

import (
    "github.com/hashicorp/go-plugin"
    "fmt"
    "os"
    "github.com/hashicorp/go-hclog"
    "github.com/go-errors/errors"
    "plugins/shared"
)

var logger hclog.Logger


type AvrPlatform struct {
}

func (p *AvrPlatform) CreateExampleProject(information *shared.ProjectInformation) ([]shared.ExecutableInformation, error) {
    fmt.Fprintln(os.Stdout, "Populated Example Project")
    return nil, nil
}

func (p *AvrPlatform) BuildProject(information *shared.TargetInformation) ([]shared.ExecutableInformation, error) {
    logger.Info(information.Platform)
    information.Platform = "Nola"
    logger.Info("Finished!!!!!")

    return []shared.ExecutableInformation{{
        CommandName: "Deep",
    }}, errors.New("This is stupid")
}

func (p *AvrPlatform) RunProject(information *shared.RunInformation) ([]shared.ExecutableInformation, error) {
    fmt.Fprintln(os.Stdout, "Running project")
    return nil, nil
}


func main() {
    logger = hclog.New(&hclog.LoggerOptions{
        Output: hclog.DefaultOutput,
        Level:  hclog.Trace,
        Name:   "avr",
    })

    plugin.Serve(&plugin.ServeConfig{
        HandshakeConfig: shared.Handshake,
        Plugins: map[string]plugin.Plugin{
            "avr-platform": &shared.PlatformPlugin{Impl: &AvrPlatform{}},
        },

        // A non-nil value here enables gRPC serving for this plugin...
        GRPCServer: plugin.DefaultGRPCServer,
    })
}
