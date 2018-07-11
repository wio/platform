package main

import (
    "github.com/hashicorp/go-plugin"
    "plugins/shared"
    "plugins/output"
)

type AvrPlatform struct {
}

func (p *AvrPlatform) CreateExampleProject(information *shared.ProjectInformation) ([]shared.ExecutableInformation, error) {
   output.Info("Populated Example Project")
    return nil, nil
}

func (p *AvrPlatform) BuildProject(information *shared.TargetInformation) ([]shared.ExecutableInformation, error) {
    output.Info("Building the project")
    output.Info("Platform Name: " + information.Platform)

    return []shared.ExecutableInformation{{
        CommandName: "Deep",
    }}, nil
}

func (p *AvrPlatform) RunProject(information *shared.RunInformation) ([]shared.ExecutableInformation, error) {
    output.Info("Running the project")
    return nil, nil
}

func main() {
    output.Info("Plugin Started")

    // Create a plugin to communicate
    plugin.Serve(&plugin.ServeConfig{
        HandshakeConfig: shared.Handshake,
        Plugins: map[string]plugin.Plugin{
            "platform-atmelavr": &shared.PlatformPlugin{Impl: &AvrPlatform{}},
        },

        // A non-nil value here enables gRPC serving for this plugin...
        GRPCServer: plugin.DefaultGRPCServer,
    })
}
