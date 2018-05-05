[![Build Status](https://travis-ci.org/mhmoudgmal/pdcli.svg?branch=master)](https://travis-ci.org/mhmoudgmal/pdcli)

PagerDuty CLI **(Under development)**

## Development
- clone the project `git clone git@github.com:mhmoudgmal/pdcli.git`
- install go dep tool `go get -u github.com/golang/dep/cmd/dep`
- run `dep ensure`

## Tests
This project uses ginkgo BDD framework https://github.com/onsi/ginkgo

- run `go test ./...`
- or install ginkgo and run `ginkgo -r`

## Run the app
- `go build`
- `./pdcli -email=your_email@registerd.pd -token=pd_token`
