# loggo [![Build Status](https://travis-ci.org/claudetech/loggo.svg?branch=master)](https://travis-ci.org/claudetech/loggo)

An easy to use and configurable logger for Go.

## Basic usage

`loggo/default` exposes a `Log` instance to use for
really simple cases.
It logs levels `<= Info` to `stdout` and level `>= Warning`
to stderr.

```go
package main

import (
  "github.com/claudetech/loggo"
  . "github.com/claudetech/loggo/default"
)

func main() {
  Log.SetLevel(loggo.Verbose)
  Log.Verbose("foobar")
  Log.Debug("foobar")
  Log.Info("foobar")
  Log.Warning("foobar")
  Log.Error("foobar")
  Log.Critical("foobar")
}
```

will give the following output

![loggo-output](http://res.cloudinary.com/dtdu3sqtl/image/upload/v1413133002/loggo_tjrayh.png)
