# Go Install

[Go Install](https://goinstall.app/) is an on-demand binary server, allowing non-Go users to quickly install tools written in Go without installing `go` itself, and removing the burden of cross-compiling and uploading dozens of binaries for every command-line you write.

![On-demand Go binary example](assets/example.png)

> Note: This is a fork of the original [tj/gobinaries](https://github.com/tj/gobinaries) project.

## Usage

Install `PKG` with optional semver `VERSION`.

```
curl -sf https://goinstall.app/<PKG>[@VERSION] | sh
```

Install `PKG` with optional semver `VERSION` to the given directory `PREFIX`. By default `/usr/local/bin` is used.

```
curl -sf https://goinstall.app/<PKG>[@VERSION] | PREFIX=/tmp sh
```

The `github.com` path prefix is optional.

## Examples

Install the `hey` HTTP benchmarking tool:

```sh
curl -sf https://goinstall.app/rakyll/hey | sh
```

Install the `wuzz` HTTP client:

```sh
curl -sf https://goinstall.app/asciimoo/wuzz | sh
```

Install the latest version of `staticgen`:

```
curl -sf https://goinstall.app/tj/staticgen/cmd/staticgen | sh
```

Install the latest `1.x` major version of `staticgen`:

```
curl -sf https://goinstall.app/tj/staticgen/cmd/staticgen@1.x | sh
```

Install the `v1.0.0` version of `triage`:

```
curl -sf https://goinstall.app/tj/triage/cmd/triage@1.0.0 | sh
```

## Semver support

The following semver patterns are supported:

- Exact version: `v1.2.3`, `1.2.3`
- Wildcards: `v1.x`, `v1.2.x`, `1.x`, `1.2.x`
- Wildcards (implicit): `v1`, `v1.2`, `1.2`
- Leading `v` is optional, regardless of the Git tag

## How does it work?

Take for example you want to install @rakyll's fantastic HTTP benchmarking tool [hey](https://goinstall.app/rakyll/hey), to install the latest release you would run:

```sh
curl -sf https://goinstall.app/rakyll/hey | sh
```

This request fetches the latest tag from GitHub and responds with a shell script
which performs a second request, populated with the resolved version and architecture as shown here:

```
https://goinstall.app/binary/github.com/rakyll/hey?os=darwin&arch=amd64&version=v0.1.3
```

The response of this request is a Golang binary compiled for the requested os, architecture, and package version. The result is cached in a CDN for subsequent requests.

## Limitations

- The Go package must compile in under 100 seconds (CDN limitation)
- The Go package must have at least one Git tag

## FAQ

<details>
  <summary>What's wrong with go-get?</summary>
  Using go-get for installation is fine if you're a Go user, however, if you're creating programs which are designed to be used by a wider audience it's unlikely they will have Go installed.
</details>

<details>
  <summary>Which version of Go is used?</summary>
  Currently Go 1.13.x via the official golang:1.13 Docker image.
</details>

---

[![GoDoc](https://godoc.org/github.com/matsilva/gobinaries?status.svg)](https://godoc.org/github.com/matsilva/gobinaries)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-maintained-green.svg)
![](https://github.com/matsilva/gobinaries/workflows/Tests/badge.svg)

## Sponsors

This project was originally sponsored by [CTO.ai](https://cto.ai/).

[![](https://apex-software.imgix.net/github/sponsors/cto.png)](https://cto.ai/)

And my [GitHub sponsors](https://github.com/sponsors/tj):

[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/0" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/0)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/1" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/1)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/2" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/2)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/3" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/3)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/4" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/4)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/5" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/5)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/6" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/6)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/7" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/7)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/8" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/8)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/9" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/9)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/10" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/10)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/11" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/11)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/12" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/12)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/13" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/13)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/14" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/14)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/15" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/15)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/16" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/16)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/17" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/17)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/18" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/18)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/19" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/19)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/20" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/20)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/21" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/21)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/22" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/22)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/23" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/23)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/24" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/24)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/25" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/25)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/26" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/26)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/27" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/27)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/28" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/28)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/29" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/29)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/30" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/30)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/31" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/31)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/32" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/32)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/33" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/33)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/34" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/34)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/35" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/35)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/36" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/36)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/37" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/37)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/38" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/38)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/39" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/39)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/40" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/40)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/41" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/41)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/42" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/42)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/43" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/43)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/44" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/44)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/45" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/45)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/46" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/46)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/47" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/47)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/48" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/48)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/49" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/49)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/50" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/50)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/51" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/51)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/52" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/52)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/53" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/53)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/54" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/54)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/55" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/55)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/56" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/56)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/57" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/57)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/58" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/58)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/59" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/59)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/60" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/60)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/61" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/61)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/62" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/62)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/63" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/63)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/64" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/64)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/65" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/65)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/66" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/66)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/67" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/67)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/68" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/68)
