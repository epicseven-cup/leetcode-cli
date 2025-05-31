# go-cli-template
template for go cli setup for github

# Go Cli install path
The path of `go install <github path>` should be the url endpoints of the github project,

```bash
go install github.com/<username>/<project>
```
Please use the readme example below as an example


# Go mod init
Important note is that when you run `go mod init`, you will need to use lthe github url endpoint for the project as pointed out in the selection above.

For example, for this project the github url endpoint will be `github.com/epicseven-cup/go-cli-template`

So when you will need to run `go mod init github.com/epicseven-cup/go-cli-template`


# Binary download

When you run `go install github.com/epicseven-cup/go-cli-template@latest`, but default it will go by the main.go file that is located in the root of the repository. 

You can also point the location to another main.go file by changing the file path examples if you have `cmd/go-cli-template/main.go` you can use:

```bash
go install github.com/epicseven-cup/go-cli-template/cmd/go-cli-template@latest
```

# Install
1. You will need to install go on your machine: https://go.dev/doc/install
2. Setup GOPATH

Add the following to your shell config
```bash
export PATH=${PATH}:$HOME/go/bin
```
More information: https://go.dev/wiki/GOPATH#gopath-variable

3. Install the binary
```bash
go install github.com/epicseven-cup/go-cli-template@latest 
```

There could be delays between the Goproxy and GitHub binarys, you can use the direct setup
```bash
GOPROXY=direct go install github.com/epicseven-cup/go-cli-template@latest
```
