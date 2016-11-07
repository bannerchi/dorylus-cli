# dorylus-cli

## Install

### use [glide](https://github.com/Masterminds/glide)

Make sure your PATH includes the $GOPATH/bin directory so your commands can be easily used:

```bash
export PATH=$PATH:$GOPATH/bin
```

```bash
go get github.com/bannerchi/dorylus-cli
glide install
go install
```


## Useage


```bash
$dorylus-cli
NAME:
   dorylus - A cli for dorylus

USAGE:
   dorylus-cli command tools for dorylus

VERSION:
   0.0.1

AUTHOR:
   Ron Chi <bannerchi@gmail.com>

COMMANDS:
     start, s   start a cron
     remove, r  remove a cron
     rtr, t     get ready to run cron job
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -V  print only the version

```