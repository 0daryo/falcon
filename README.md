# Falcon

golang template for web api

## Directory

```
.
├── LICENSE
├── README.md
├── api
├── build
│   └── server
├── cmd
│   └── server
├── config
├── database
├── go.mod
├── go.sum
├── internal
│   └── server
├── main
└── pkg
    └── lib
```

inspired by https://github.com/golang-standards/project-layout

## Library

- DB: mysql
- ORM: github.com/jinzhu/gorm
- realize as hotreload tool

## To produce pb

In the directory of proto,
`protoc -I . user.proto --go_out=plugins=grpc:../../pb/server`
