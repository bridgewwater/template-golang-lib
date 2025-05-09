## install github app

- [https://github.com/apps/codecov/installations/new](https://github.com/apps/codecov/installations/new)
- add `CODECOV_TOKEN` to action secrets

## config-file

`golang-codecov.yml`

```yml
name: golang-codecov

on:
  push:
    tags:
      - '*' # Push events to matching *, i.e. 1.0.0 v1.0, v20.15.10

permissions:
  contents: write

jobs:
  golang-codecov:
    name: golang-codecov
    strategy:
      matrix:
        go:
          - '^1.23'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Set up Go SDK
        uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go }}
          cache: false
      - name: Print env info
        run: |
          go env
          go version

      - name: Run go build
        run: go build -v ./...

      - name: Run test coverage
        run: go test -cover -coverprofile coverage.txt -covermode count -coverpkg ./... -tags test -v ./...

      - name: Codecov
        uses: codecov/codecov-action@v5.4.2
        with:
          token: ${{secrets.CODECOV_TOKEN}}
          files: coverage.txt
          dry_run:  true
#          verbose: true

```