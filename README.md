# Leetcode-cli
cli tool to get leetcode daily question

# Example Usage:
```bash
leetcode-cli daily go
```

```bash
leetcode-cli daily python
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
go install github.com/epicseven-cup/leetcode-cli@latest
```

There could be delays between the Goproxy and GitHub binarys, you can use the direct setup
```bash
GOPROXY=direct go install github.com/epicseven-cup/leetcode-cli@latest
```
