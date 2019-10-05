# Spectre Go Rest API



### Prerequisites

| Item        | Link           | Terminal  |
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

### Dependencies

To install any of this packages, open terminal and run
`go get github.com/{package-name}`
This command will install the packages into your GOPATH.

| Item        | Link           | Terminal |
| ------------- |:-------------| -----:|
| gorilla/mux | https://github.com/gorilla/mux | `go get github.com/gorilla/mux` |
| jinzhu/gorm | https://github.com/jinzhu/gorm |   `go get github.com/jinzhu/gorm`|
| dgrijalva/jwt-go | https://github.com/dgrijalva/jwt-go | `go get github.com/dgrijalva/jwt-go` |
|joho/godotenv| https://github.com/joho/godotenv |`go get github.com/joho/godotenv` |





### References
- https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
- https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da