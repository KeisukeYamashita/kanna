# Kanna

> Kanna is a light webhook test server for webhook application developing

[![Go][go-badge]][go]
![CircleCI](https://circleci.com/gh/KeisukeYamashita/kanna.svg?style=svg&circle-token=e4b3002b1fb96c423ed5f75332c3455d88d56b0f)
[![GolangCI][golangci-badge]][golangci]
[![Go Report Card][go-report-card-badge]][go-report-card]
[![GoDoc][godoc-badge]][godoc]
[![Dependabot][dependabot-badge]][dependabot]
![License](https://img.shields.io/badge/license-Apache%202.0-%23E93424)

## How to run

See how to use Kanna.

### Basic

The request will be in stdout.

```
$ go run cmd/kanna/kanna.go
```

### Flags

* `--type`: You can select `stdout`, `json`.
* `--statusCode`: You can select all status code. 

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)

## License

Copyright 2019 The Kanna Authors. 

Kanna is released under the Apache License 2.0.

<!-- badge links -->
[go]: https://golang.org/dl
[go-badge]: https://img.shields.io/badge/Go-1.13-blue

[godoc]: https://godoc.org/github.com/KeisukeYamashita/kanna
[godoc-badge]: https://img.shields.io/badge/godoc.org-reference-blue.svg

[golangci]: https://golangci.com/r/github.com/KeisukeYamashita/kanna
[golangci-badge]: https://golangci.com/badges/github.com/KeisukeYamashita/kanna.svg

[go-report-card]: https://goreportcard.com/report/github.com/KeisukeYamashita/kanna
[go-report-card-badge]: https://goreportcard.com/badge/github.com/KeisukeYamashita/kanna

[dependabot]: https://dependabot.com 
[dependabot-badge]: https://badgen.net/badge/icon/Dependabot?icon=dependabot&label&color=blue