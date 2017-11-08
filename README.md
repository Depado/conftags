# conftags

![Go Version](https://img.shields.io/badge/go-1.8-brightgreen.svg)
![Go Version](https://img.shields.io/badge/go-1.9-brightgreen.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Depado/conftags)](https://goreportcard.com/report/github.com/Depado/conftags)
[![Build Status](https://drone.depado.eu/api/badges/Depado/conftags/status.svg)](https://drone.depado.eu/Depado/conftags)
[![codecov](https://codecov.io/gh/Depado/conftags/branch/master/graph/badge.svg)](https://codecov.io/gh/Depado/conftags)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Depado/bfchroma/blob/master/LICENSE)

Similar to [github.com/caarlos0/env](https://github.com/caarlos0/env) but with
default decoupled from environment. It also doesn't support slices.

This library intends to work as the last step in the filling of a configuration
struct. Assuming that environment variables are the highest priority
configuration values, you could first parse a configuration file, and then parse
this struct with this library. The last resort is the default which will fill
the struct field only if it has the zero value for its own type.

## Usage

```go
type MyConf struct {
    Int      int           `env:"INT" default:"10"`
    String   string        `env:"STRING"`
    Duration time.Duration `env:"DURATION" default:"1h"`
}

func (ms *MyConf) Parse() error {
    return conftags.Parse(ms)
}
```
