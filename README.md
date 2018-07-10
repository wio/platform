# Platform Plugin

This is go package that must be used to create a platform plugin so that wio can include this at run time.
This plugin will communicate over GRPC and each platform will perform their own designed tasks.

To use this package
```golang
import "github.com/wio/platform-plugin/shared"
```

## Interface
This package defines a generic interface that each platform must use. Failing to do that will result in platform being invalid.
Interface defined is:
```golang
// Greeter is the interface that we're exposing as a plugin.
type Platform interface {
    // This will create an example project and provides executables to execute
    CreateExampleProject(*ProjectInformation) ([]ExecutableInformation, error)

    // Creates build files for the project and provides executables to execute
    BuildProject(*TargetInformation) ([]ExecutableInformation, error)

    // Creates hardware specific files and provides executables to execute
    RunProject(information *RunInformation) ([]ExecutableInformation, error)
}
```

Essentially each plugin will hande 3 things:
* Creating an example project
* Creating build files for the project
* Executing the project

Execution of external commands is handled by wio and no external command must be called by a platform. Hence, everytime a function is called for a platform, it can return a list of executable information that can be used by wio to execute commands.

## Example
For a detailed example, check [tests/avr-platform directory](https://github.com/wio/platform-plugin/tree/master/tests/avr-platform)
