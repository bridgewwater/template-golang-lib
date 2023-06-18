[![golang-ci](https://github.com/bridgewwater/template-golang-lib/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/bridgewwater/template-golang-lib/actions/workflows/golang-ci.yml)
[![TravisBuildStatus](https://api.travis-ci.com/bridgewwater/template-golang-lib.svg?branch=main)](https://travis-ci.com/bridgewwater/template-golang-lib)
[![license](https://img.shields.io/github/license/bridgewwater/template-golang-lib)](https://github.com/bridgewwater/template-golang-lib)
[![go mod version](https://img.shields.io/github/go-mod/go-version/bridgewwater/template-golang-lib?label=go.mod)](https://github.com/bridgewwater/template-golang-lib)
[![GoDoc](https://godoc.org/github.com/bridgewwater/template-golang-lib?status.png)](https://godoc.org/github.com/bridgewwater/template-golang-lib/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bridgewwater/template-golang-lib)](https://goreportcard.com/report/github.com/bridgewwater/template-golang-lib)
[![codecov](https://codecov.io/gh/bridgewwater/template-golang-lib/branch/main/graph/badge.svg)](https://codecov.io/gh/bridgewwater/template-golang-lib)
[![github release](https://img.shields.io/github/v/release/bridgewwater/template-golang-lib?style=social)](https://github.com/bridgewwater/template-golang-lib/releases)

## for what

- this project used to github golang lib project

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/bridgewwater/template-golang-lib.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/bridgewwater/template-golang-lib
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/bridgewwater/template-golang-lib | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## evn

- minimum go version: go 1.17

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |

## usage

- use to replace
  `bridgewwater/template-golang-lib` to you code

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```

### cli tools to init project fast

```bash
$ curl -L --fail https://raw.githubusercontent.com/bridgewwater/template-golang-lib/main/template-golang-lib
# let template-golang-lib file folder under $PATH
$ chmod +x template-golang-lib
# see how to use
$ template-golang-lib -h
```