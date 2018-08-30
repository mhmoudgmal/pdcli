[![Build Status](https://travis-ci.org/mhmoudgmal/PD.cli.svg?branch=master)](https://travis-ci.org/mhmoudgmal/PD.cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/mhmoudgmal/PD.cli)](https://goreportcard.com/report/github.com/mhmoudgmal/PD.cli)
[![Generic badge](https://img.shields.io/badge/editor-vim-yellowgreen.svg)](https://github.com/mhmoudgmal/PD.cli)

PagerDuty CLI **(Under development)**

## Development
- clone the project `git clone git@github.com:mhmoudgmal/PD.cli.git pdcli`
- install go dep tool `go get -u github.com/golang/dep/cmd/dep`
- run `dep ensure`

## Tests
This project uses ginkgo BDD framework https://github.com/onsi/ginkgo

- run `go test ./...`
- or install ginkgo and run `ginkgo -r`

## Run the app
- `go build`
- `./pdcli -email=your_email@registerd.pd -token=pd_token`
