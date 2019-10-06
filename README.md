# Spectre Go Rest API

## Prerequisites

| Item        | Link           | Terminal  |
| ------------- |:-------------| -----:|
| Go      | <https://golang.org> | `brew install golang` |
| Visual Studio Code      | <https://code.visualstudio.com>      |   N/A |
| VS Code Go Extension | <https://code.visualstudio.com/docs/languages/go>     |   N/A |

## Setting `GOPATH`

The `GOPATH` environment variable specifies the location of your go workspace.

```bash
# First thing first, create the go workspace under your `Developer` folder.
$ mkdir -p $HOME/Developer/go

# Update go environment value of key GOPATH
$ go env -w GOPATH=$HOME/Developer/go
```

Now setup your `bash` or `zsh` accordingly, following these guides:
<https://github.com/golang/go/wiki/SettingGOPATH>

## Building the product

This is how we will be building and running our API:

```bash
# Create the executable and store it to the named output file
$ go build -o ./.build/spectre-api

# Run the exectuable
$ .build/./spectre-api
```

Additionally you can debug the current solution right from witin Visual Studio code. In order to be able to run the solution pressing the `F5` key (use `fn` key to load the function keys on the touch bar), you have to install the `dlv` package first.

PRessing `F5` without the `dlv` package installed will prompt you with this error:

> The "dlv" command is not available. Run "go get -v github.com/go-delve/delve/cmd/dlv" to install.

Go ahead and press the fix button in the editor or run the command suggested above in terminal.

## Dependencies

To install any of this packages, open terminal and run
`go get github.com/{package-name}`
This command will install the packages into your GOPATH.

| Item        | Link           | Terminal |
| ------------- |:-------------| -----:|
| gorilla/mux | <https://github.com/gorilla/mux> | `go get github.com/gorilla/mux` |
| jinzhu/gorm | <https://github.com/jinzhu/gorm> |   `go get github.com/jinzhu/gorm`|
| dgrijalva/jwt-go | <https://github.com/dgrijalva/jwt-go> | `go get github.com/dgrijalva/jwt-go` |
| joho/godotenv| <https://github.com/joho/godotenv> |`go get github.com/joho/godotenv` |
| golang/protobuf| <https://github.com/golang/protobuf> | `go get github.com/golang/protobuf/proto` |

## References

- <https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b>
- <https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da>
