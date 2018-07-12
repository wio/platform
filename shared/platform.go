package shared

import (
    "github.com/hashicorp/go-plugin"
    "google.golang.org/grpc"
    "platform-plugin/proto"
    "context"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
    ProtocolVersion:  1,
    MagicCookieKey:   "PLATFORM_PLUGIN",
    MagicCookieValue: "wio",
}

// Platform is the interface that we're exposing as a plugin.
type Platform interface {
    // This will create an example project and provides executables to execute
    CreateExampleProject(*ProjectInformation) ([]ExecutableInformation, error)

    // Creates build files for the project and provides executables to execute
    BuildProject(*TargetInformation) ([]ExecutableInformation, error)

    // Creates hardware specific files and provides executables to execute
    RunProject(information *RunInformation) ([]ExecutableInformation, error)
}


type PlatformPlugin struct {
    plugin.NetRPCUnsupportedPlugin
    Impl Platform
}

func (p *PlatformPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
    proto.RegisterPlatformServer(s, &GRPCServer{
        Impl:   p.Impl,
        broker: broker,
    })
    return nil
}

func (p *PlatformPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
    return &GRPCClient{
        client: proto.NewPlatformClient(c),
        broker: broker,
    }, nil
}
var _ plugin.GRPCPlugin = &PlatformPlugin{}
