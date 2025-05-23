# dev

## just start dev

- minimum go version: go 1.23
- change `go 1.23`, `^1.23`, `1.23.8` to new go version
- change `golangci-lint@v2.1.6` from [golangci-lint version release](https://github.com/golangci/golangci-lint/releases) to new version
    - more info see [golangci-lint installation](https://golangci-lint.run/welcome/install/)

## dev tasks

```bash
# It needs to be executed after the first use or update of dependencies.
make init dep
```

- test code

```bash
make test
# benchmark and coverage show
make ci.test.benchmark ci.coverage.show
```

- ci to fast check as CI pipeline

```bash
# check style at local
make style

# run ci at local
make ci
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
