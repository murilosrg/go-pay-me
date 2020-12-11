# PayMe API

[![codecov](https://codecov.io/gh/murilosrg/go-pay-me/branch/main/graph/badge.svg?token=LgWmKSXTOz)](https://codecov.io/gh/murilosrg/go-pay-me)
[![go-report](https://goreportcard.com/badge/github.com/murilosrg/go-pay-me?style=flat-square)](https://goreportcard.com/report/github.com/murilosrg/go-pay-me)
[![go-mod](https://img.shields.io/github/go-mod/go-version/murilosrg/go-pay-me?style=flat-square)](https://github.com/murilosrg/go-pay-me)
[![license](https://img.shields.io/github/license/murilosrg/go-pay-me?style=flat-square)](https://opensource.org/licenses/MIT)

🖖🏻 A simple api application written in Go

## Quick Start (Docker Deploy)

```sh
$ docker-compose up --build
```

## Development (Non-Dockerized Deploy)

### 1.Clone Source Code

```shell
$ git clone https://github.com/murilosrg/go-pay-me

$ cd go-pay-me
```

### 2.Download Requirements

```shell
$ go mod download
```

### 3.Create Configuration

```shell
$ touch /etc/payme/configuration.yaml
```

```yaml
# configuration for dev
db:
  driver: sqlite3
  addr: ./payme.db
address: ":8080"
acquires:
  stoneUrl: "http://demo4605147.mockable.io/stone/"
  cieloUrl: "http://demo4605147.mockable.io/cielo/"
```

| Param     | Description                                           | 
| --------- | ----------------------------------------------------- | 
| db        | Database configure, supports sqlite3, mysql, postgres | 
| address   | Listen address                                        | 
| acquires  | url of acquires mocked                                | 


### 4.Init and Run

```shell
$ (sudo) go run ./cmd/payme -init
```

## Run Test

```shell
$ (sudo) go test ./cmd/payme
```

## License

[MIT](https://opensource.org/licenses/MIT)