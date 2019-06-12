# Simple Matrix client in go

This project uses the [matrix](https://matrix.org/) APIs to connect to a specific room and to read and post messages.

## How to use
If you want to use it with your own instance of riot, you have to change few things:
    - create a file "riot.pass" with your user password
    - change the user in the main function
    - change the const BaseURL to point to your server
    - change the const RoomID to point to your room

## Dev
This go program is using the go modules:
- remove the gopath : `export GOPATH=""`
- switch on the go modules: `export GO111MODULE=on"`
- inside the empty project folder, create the project with `go mod init project_name`
- add resty dep with `go get gopkg.in/resty.v1`
- add fastjson dep with `go get github.com/valyala/fastjson`