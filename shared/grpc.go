package shared

import (
    "github.com/hashicorp/go-plugin"
    "plugins/proto"
    "context"
    "errors"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
    broker *plugin.GRPCBroker
    client proto.PlatformClient
}

func protoExecutablesToExecutableInformation(executables *proto.Executables) []ExecutableInformation {
    if executables == nil || executables.Executables ==  nil {
        return nil
    }

    var execInformationToReturn []ExecutableInformation

    for _, executable := range executables.Executables {
        execInformationToReturn = append(execInformationToReturn, ExecutableInformation{
            CommandName: executable.CommandName,
            CommandArguments: executable.CommandArguments,
        })
    }

    return execInformationToReturn
}

func executableInformationToProtoExecutables(executablesInformation []ExecutableInformation) *proto.Executables {
    if executablesInformation == nil {
        return nil
    }

    executables := &proto.Executables{}

    for _, executable := range executablesInformation {
        executables.Executables = append(executables.Executables, &proto.Executable{
            CommandName: executable.CommandName,
            CommandArguments: executable.CommandArguments,
        })
    }

    return executables
}

func (m *GRPCClient) CreateExampleProject(projectInformation *ProjectInformation) ([]ExecutableInformation, error) {
    if projectInformation == nil {
        return nil, errors.New("project information cannot be nil")
    }

    executables, err := m.client.CreateExampleProject(context.Background(), &proto.ProjectInformation{
        Name: projectInformation.Name,
        Directory: projectInformation.Directory,
        WioPath: projectInformation.WioPath,
        ProjectType: projectInformation.ProjectType,
    })

    return protoExecutablesToExecutableInformation(executables), err
}

func (m *GRPCClient) BuildProject(targetInformation *TargetInformation) ([]ExecutableInformation, error) {
    if targetInformation == nil {
        return nil, errors.New("target information cannot be nil")
    }

    targetDependencies := map[string]*proto.DependencyTarget{}
    for dependencyName, dependency := range targetInformation.Dependencies {
        targetDependencies[dependencyName] = &proto.DependencyTarget{
            Name: dependency.Name,
            Path: dependency.Path,
            Flags: dependency.Flags,
            Definitions: dependency.Definitions,
            FlagsVisibility: dependency.FlagsVisibility,
            DefinitionsVisibility: dependency.DefinitionsVisibility,
            HeaderOnly: dependency.HeaderOnly,
        }
    }

    var targetLinkDependencies []*proto.DependencyLinkTarget
    for _, dependencyLink := range targetInformation.DependenciesLink {
        targetLinkDependencies = append(targetLinkDependencies, &proto.DependencyLinkTarget{
            From: dependencyLink.From,
            To: dependencyLink.To,
            Visibility: dependencyLink.Visibility,
        })
    }

    executables, err := m.client.BuildProject(context.Background(), &proto.TargetInformation{
        ProjectInformation: &proto.ProjectInformation{
            Name: targetInformation.ProjectInformation.Name,
            Directory: targetInformation.ProjectInformation.Directory,
            WioPath: targetInformation.ProjectInformation.WioPath,
            ProjectType: targetInformation.ProjectInformation.ProjectType,
        },
        TargetName: targetInformation.TargetName,
        TargetPath: targetInformation.TargetPath,
        Platform: targetInformation.Platform,
        Framework: targetInformation.Framework,
        Hardware: targetInformation.Hardware,
        Flags: targetInformation.Flags,
        Definitions: targetInformation.Definition,
        Dependencies: targetDependencies,
        DependenciesLink: targetLinkDependencies,
    })

    return protoExecutablesToExecutableInformation(executables), err
}

func (m *GRPCClient) RunProject(uploadInformation *RunInformation) ([]ExecutableInformation, error) {
    if uploadInformation == nil {
        return nil, errors.New("upload information cannot be nil")
    }

    executables, err := m.client.RunProject(context.Background(), &proto.RunInformation{
        TargetName: uploadInformation.TargetName,
        HardwarePort: uploadInformation.HardwarePort,
    })

    return protoExecutablesToExecutableInformation(executables), err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
    // This is the real implementation
    Impl Platform

    broker *plugin.GRPCBroker
}

func (m *GRPCServer) CreateExampleProject(ctx context.Context, args *proto.ProjectInformation) (*proto.Executables, error) {
    if args == nil {
        return nil, errors.New("project information cannot be nil")
    }

    projectInformation := &ProjectInformation{
        Name: args.Name,
        Directory: args.Directory,
        WioPath: args.WioPath,
        ProjectType: args.ProjectType,
    }

    executableInformation, err := m.Impl.CreateExampleProject(projectInformation)
    return executableInformationToProtoExecutables(executableInformation), err
}

func (m *GRPCServer) BuildProject(ctx context.Context, args *proto.TargetInformation) (*proto.Executables, error) {
    if args == nil {
        return nil, errors.New("target information cannot be nil")
    }

    targetDependencies := map[string]*DependencyTarget{}
    for dependencyName, dependency := range args.Dependencies {
        targetDependencies[dependencyName] = &DependencyTarget{
            Name: dependency.Name,
            Path: dependency.Path,
            Flags: dependency.Flags,
            Definitions: dependency.Definitions,
            FlagsVisibility: dependency.FlagsVisibility,
            DefinitionsVisibility: dependency.DefinitionsVisibility,
            HeaderOnly: dependency.HeaderOnly,
        }
    }

    var targetLinkDependencies []*DependencyLinkTarget
    for _, dependencyLink := range args.DependenciesLink {
        targetLinkDependencies = append(targetLinkDependencies, &DependencyLinkTarget{
            From: dependencyLink.From,
            To: dependencyLink.To,
            Visibility: dependencyLink.Visibility,
        })
    }

    targetInformation := &TargetInformation{
        ProjectInformation: ProjectInformation{
            Name: args.ProjectInformation.Name,
            Directory: args.ProjectInformation.Directory,
            WioPath: args.ProjectInformation.WioPath,
            ProjectType: args.ProjectInformation.ProjectType,
        },
        TargetName: args.TargetName,
        TargetPath: args.TargetPath,
        Platform: args.Platform,
        Framework: args.Framework,
        Hardware: args.Hardware,
        Flags: args.Flags,
        Definition: args.Definitions,
        Dependencies: targetDependencies,
        DependenciesLink: targetLinkDependencies,
    }

    executableInformation, err := m.Impl.BuildProject(targetInformation)
    return executableInformationToProtoExecutables(executableInformation), err
}

func (m *GRPCServer) RunProject(ctx context.Context, args *proto.RunInformation) (*proto.Executables, error) {
    if args == nil {
        return nil, errors.New("upload information cannot be nil")
    }

    executableInformation, err := m.Impl.RunProject(&RunInformation{
        TargetName: args.TargetName,
        HardwarePort: args.HardwarePort,
    })

    return executableInformationToProtoExecutables(executableInformation), err
}
