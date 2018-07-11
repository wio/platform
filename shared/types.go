package shared

// Information about wio executable
type ExecutableInformation struct {
    CommandName string
    CommandArguments []string
}

// Information about the project for which wio is being used
type ProjectInformation struct {
    Name string
    Directory string
    WioPath string
    ProjectType string
    HeaderOnly bool
    Framework string
}

// Dependency of a project and it's information
type DependencyTarget struct {
    Name            string
    Path                  string
    Flags                 []string
    Definitions           []string
    FlagsVisibility       string
    DefinitionsVisibility string
    HeaderOnly            bool
}

// Link target to link dependencies
type DependencyLinkTarget struct {
    From       string
    To         string
    Visibility string
}

// Information about target being built
type TargetInformation struct {
    ProjectInformation ProjectInformation
    TargetName string
    TargetPath string
    Platform string
    Framework string
    Hardware string
    Flags []string
    Definition [] string

    Dependencies map[string]*DependencyTarget
    DependenciesLink []*DependencyLinkTarget
}

// Information about executing a target
type RunInformation struct {
    TargetName string
    HardwarePort string
}
