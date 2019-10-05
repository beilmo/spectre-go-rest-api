# Spectre Go Rest API



### Prerequisites

| Item        | Link           | Brew  |
| ------------- |:-------------| -----:|
| Go      | https://golang.org | `brew install golang` |
| Visual Studio Code      | https://code.visualstudio.com      |   N/A |
| VS Code Go Extension | https://code.visualstudio.com/docs/languages/go     |   N/A |



### Setting `GOPATH`

The `GOPATH` environment variable specifies the location of your go workspace.

```bash
# First thing first, create the go workspace under your `Developer` folder.
$ mkdir -p $HOME/Developer/go

# Update go environment value of key GOPATH
$ go env -w GOPATH=$HOME/Developer/go
```
Now setup your `bash` or `zsh` accordingly, following these guides:
https://github.com/golang/go/wiki/SettingGOPATH



### Building the product

This is how we will be building and running our API:
```bash
# Create the executable and store it to the named output file 
$ go build -o ./.build/spectre-api

# Run the exectuable
$ .build/./spectre-api
```
