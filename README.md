# go-mpd

Go package to parse, manipulate and build MPEG-DASH (ISO/IEC 23009-1 5th edition) manifests.

[![Build Status](https://github.com/joeig/go-mpd/workflows/Tests/badge.svg)](https://github.com/joeig/go-mpd/actions)
[![Test coverage](https://img.shields.io/badge/coverage-92%25-success)](https://github.com/joeig/go-mpd/tree/master/.github/testcoverage.yml)
[![Go Report Card](https://goreportcard.com/badge/go.eigsys.de/go-mpd)](https://goreportcard.com/report/go.eigsys.de/go-mpd)
[![PkgGoDev](https://pkg.go.dev/badge/go.eigsys.de/go-mpd)](https://pkg.go.dev/go.eigsys.de/go-mpd)

## Usage

```shell
go get -u go.eigsys.de/go-mpd
```

### Forked version of `encoding/xml`

This package is using a forked version of `encoding/xml` by [ydnar](https://github.com/ydnar/go/tree/xmlns-fixes)
which adds support for XMLNS prefixes.
It is planned to move back to the standard library once it supports XMLNS prefixes.
([#48641](https://github.com/golang/go/pull/48641)).

Please note that the forked `encoding/xml` package is still under active development.
It has not been approved by the Go maintainers yet
and therefore may not meet the quality standards of the standard library.

## Examples

A complete list of examples is available in the [package reference](https://pkg.go.dev/go.eigsys.de/go-mpd).

### Create empty MPD

```go
package main

import (
	"go.eigsys.de/go-mpd"
)

func main() {
	example := mpd.New()
	example.MinBufferTime = "PT2S"
	example.Period = []mpd.Period{{ID: "period-0"}}
}

```

### Read MPD from file

```go
package main

import (
	"go.eigsys.de/go-mpd"
	"log"
	"os"
)

func main() {
	handle, err := os.Open("manifest.mpd")
	if err != nil {
		log.Fatalf("%v", err)
	}

	example, err = mpd.Read(handle)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

```

### Marshal MPD

```go
package main

import (
	"go.eigsys.de/go-mpd"
	"log"
)

func main() {
	example := mpd.New()
	example.Profiles = mpd.ISOFFOnDemand2011Profile
	example.Type = mpd.StaticPresentationType
	example.MinBufferTime = "PT2S"
	example.Period = []mpd.Period{{ID: "period-0"}}

	exampleBytes, err := example.Bytes()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

```
