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

## Things to Note
### Handle only 3 major things
Essentially each plugin will hande 3 things:
* Creating an example project
* Creating build files for the project
* Executing the project

### Command Execution
Execution of external commands is handled by wio and no external command must be called by a platform. Hence, everytime a function is called for a platform, it can return a list of executable information that can be used by wio to execute commands.

### Logs
If a plugin wants to print anything, you will have to use a provided `logger`. If you do not use it, your output will not be registered. To use logger:
```golang
import "github.com/wio/platform-plugin/output"

// Logger function
output.Info("Hi")
output.Warn("Warning")
output.Fatal("Error")
ouput.Verb("Verbose")

// Regular ouput
output.Print("No new line")
output.Println("new line")
output.Printf("Formatted line %s", "here")
```

### Platform.json
Each plugin must have a `platform.json` file. This file provides information about the plugin and essentially helps `wio` determine configurations.

```json
{
  "name": "platform-atmelavr",
  "exec": "avr"
}
```
* Make sure `exec` tag is properly set to the name of the executable

### Creating a server
Plugin is a server to wio which is a client. This means you will have to add some boiler code to make sure plugin is properly defined.
```golang
// Create a plugin to communicate
    plugin.Serve(&plugin.ServeConfig{
        HandshakeConfig: shared.Handshake,
        Plugins: map[string]plugin.Plugin{
            "platform-name": &shared.PlatformPlugin{Impl: &PlatformStruct{}},
        },

        // A non-nil value here enables gRPC serving for this plugin...
        GRPCServer: plugin.DefaultGRPCServer,
    })
```
* This must be added to your main function
* In plugins, the name of the plugin must be the same as the one defined in `platform.json`. This is because client needs to know which plugin to call.


## Example
Example of a server looks like this:
```golang
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
    return []shared.ExecutableInformation{{
        CommandName: "Deep",
    }}, nil
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

```

For a detailed example, check [tests directory](https://github.com/wio/platform-plugin/tree/master/tests/)
